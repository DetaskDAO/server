package blockchain

import (
	ABI "code-market-admin/abi"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/utils"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"sync/atomic"
	"time"
)

var lock atomic.Bool

func BalanceRPC(chain string) {
	if lock.Load() {
		return
	}
	lock.Store(true)
	defer lock.Store(false)
	blockChain := global.CONFIG.BlockChain
	providerMap := make(map[string]string)
	for _, v := range blockChain {
		if chain != "" && chain != v.Name {
			providerMap[v.Name] = global.ProviderMap[v.Name]
			continue
		}
		if len(v.Provider) == 0 {
			return
		}
		indexList := make([]int64, len(v.Provider))
		for i, url := range v.Provider {
			spent, _ := rpcRequest(chain, url)
			indexList[i] = spent
		}
		i, _ := utils.SliceMin[int64](indexList)
		providerMap[v.Name] = v.Provider[i]
		global.LOG.Warn("RPC 切换: " + v.Name + " " + strconv.Itoa(i))
	}
	global.ProviderMap = providerMap

}

func rpcRequest(chain, url string) (spent int64, err error) {
	defer func() {
		if err := recover(); err != nil {
			spent = 9999999999999
			return
		}
	}()
	startTime := time.Now()
	rpcClient, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	blockNumber, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		panic(err)
	}
	address := global.ContractAddr[chain+":DeOrder"]
	instance, err := ABI.NewDeOrder(address, rpcClient)
	if err != nil {
		fmt.Println(err)
	}
	order, err := instance.GetOrder(nil, big.NewInt(1))
	if err != nil || order.Issuer.String() == "" || blockNumber == 0 {
		return 9999999999999, errors.New("error")
	}
	return time.Since(startTime).Milliseconds(), nil
}
