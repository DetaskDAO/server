package blockchain

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

var Traversed sync.Map // 任务运行状态

func HandleTransaction(chain string) {
	_, ok := Traversed.Load(chain)
	if ok {
		return
	}
	Traversed.Store(chain, true)
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			Traversed.Delete(chain)
			global.LOG.Error("HandleTransaction致命错误")
			go BalanceRPC(chain)
			time.Sleep(time.Second * 1)
			go HandleTransaction(chain)
		}
	}()

	client, err := ethclient.Dial(global.ProviderMap[chain])
	if err != nil {
		panic("Error dial")
	}
	var txMap sync.Map
	var countMap sync.Map
	// 循环
	for {
		// 超出扫描次数删除
		countMap.Range(func(key, value interface{}) bool {
			v, ok := value.(int)
			if ok && v > 100 {
				fmt.Println("超出扫描次数删除")
				HandleTraverseFailed(chain, key.(string), 3)
				countMap.Delete(key)
			}
			return true
		})
		// 获取需要扫描的数据
		var transHashList []model.TransHash
		db := global.MAPDB[chain].Model(&model.TransHash{})
		if err := db.Find(&transHashList).Error; err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		var haveBool bool // 是否空map
		txMap.Range(func(key, value interface{}) bool {
			haveBool = true
			return false
		})
		// 无任务
		if len(transHashList) == 0 && !haveBool {
			Traversed.Delete(chain)
			return
		}
		// 任务列表
		for _, trans := range transHashList {
			trans.Hash = strings.TrimSpace(trans.Hash)
			_, loaded := txMap.LoadOrStore(trans.Hash, trans)
			if loaded == false {
				go HandleTransactionReceipt(client, chain, &txMap, &countMap, trans.Hash)
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func HandleTransactionReceipt(client *ethclient.Client, chain string, txMap *sync.Map, countMap *sync.Map, hash string) {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			txMap.Delete(hash)
			global.LOG.Error("HandleTransactionReceipt致命错误", zap.Any("err ", err))
		}
	}()
	// 是否在处理列表
	transHashAny, ok := txMap.Load(hash)
	if !ok {
		return
	}
	transHash, ok := transHashAny.(model.TransHash)
	if !ok {
		global.LOG.Error("HandleTransactionReceipt Reflect Error")
		return
	}
	fmt.Println(hash)
	// 解析交易Hash
	res, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	// 待交易
	if err != nil {
		fmt.Println("待交易", err)
		txMap.Delete(hash)
		// TODO: 控制尝试次数
		times, exist := countMap.LoadOrStore(hash, 1)
		if exist {
			v, ok := times.(int)
			if ok {
				countMap.Store(hash, v+1)
			}
		}
		return
	}
	// 交易失败
	if res.Status == 0 {
		fmt.Println("交易失败")
		txMap.Delete(hash)
		countMap.Delete(hash)
		if err = HandleTraverseFailed(chain, transHash.Hash, 2); err != nil {
			fmt.Println(err)
		}
		return
	}
	// 交易成功
	if res.Status == 1 {
		fmt.Println("交易成功")
		if err := EventsParser(chain, transHash, res.Logs); err != nil {
			fmt.Println(err)
			txMap.Delete(hash)
		} else {
			fmt.Println("交易成功--删除")
			txMap.Delete(hash)
			countMap.Delete(hash)
		}
	}
}
