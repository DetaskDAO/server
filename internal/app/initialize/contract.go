package initialize

import (
	ABI "code-market-admin/abi"
	ABIFILE "code-market-admin/abi"
	"code-market-admin/internal/app/global"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
)

// InitContract 加载合约信息
func InitContract() {
	initContractAddr()
	initContractEvent()
}

// initContractEvent 加载合约Event信息
func initContractEvent() {
	global.ContractEvent = make(map[common.Hash]string)
	ABIList := []string{ABI.DeOrderMetaData.ABI, ABI.DeTaskMetaData.ABI}
	for _, abiStr := range ABIList {
		contractAbi, err := abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			panic(err)
		}
		for _, v := range contractAbi.Events {
			global.ContractEvent[v.ID] = v.Name
		}
	}
}

// initContractAddr 加载合约地址信息
func initContractAddr() {
	// 初始化全局变量
	global.ContractAddr = make(map[string]common.Address)
	// 遍历当前目录下的JSON文件，获取合约地址
	for _, v := range global.CONFIG.BlockChain {
		if err := fs.WalkDir(ABIFILE.ABIFS, v.Name, walkFunc); err != nil {
			panic(err)
		}
	}

}

func walkFunc(path string, info fs.DirEntry, err error) error {
	dir := filepath.Dir(path) // 目录名称
	// 判断文件是否是JSON文件
	if !info.IsDir() && filepath.Ext(path) == ".json" {
		// 打开文件
		file, err := ABIFILE.ABIFS.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 解析JSON文件
		byteValue, _ := io.ReadAll(file)
		// 结构体
		type ContractAddr struct {
			Address      string `json:"address"`
			ContractName string `json:"contractName"`
		}
		var contractAddr ContractAddr
		err = json.Unmarshal(byteValue, &contractAddr)
		if err != nil {
			fmt.Println(err)
		}
		mapName := dir + ":" + contractAddr.ContractName
		global.ContractAddr[mapName] = common.HexToAddress(contractAddr.Address)
	}

	return nil
}
