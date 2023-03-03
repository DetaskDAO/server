package initialize

import (
	"code-market-admin/internal/app/blockchain"
	"code-market-admin/internal/app/global"
)

func SweepTask() {
	for _, v := range global.CONFIG.BlockChain {
		go blockchain.HandleTransaction(v.Name)
	}
}

func ReadProvider() {
	global.ProviderMap = make(map[string]string)
	for _, v := range global.CONFIG.BlockChain {
		global.ProviderMap[v.Name] = v.Provider[0]
	}
}
