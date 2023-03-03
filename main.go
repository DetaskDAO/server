package main

import (
	"code-market-admin/internal/app/blockchain"
	"code-market-admin/internal/app/core"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/initialize"
	"code-market-admin/internal/app/timer"
	"code-market-admin/internal/app/utils"
	"go.uber.org/zap"
	"time"
)

func main() {
	global.StartTime = time.Now()
	// 初始化Viper
	core.Viper()
	// 初始化zap日志库
	global.LOG = core.Zap()
	// 注册全局logger
	zap.ReplaceGlobals(global.LOG)
	// 初始化多链 GORM连接
	initialize.InitChainDB()
	// 初始化数据库
	initialize.InitCommonDB()
	// 初始化合约地址
	initialize.InitContract()
	// 初始化缓存
	global.JsonCache = initialize.JsonCache()
	global.TokenCache = initialize.TokenCache()
	// 初始化合约内容
	go utils.BalanceIPFS()
	go blockchain.BalanceRPC("")
	initialize.ReadProvider()
	// 定时任务
	timer.Timer()
	// 启动扫块任务
	initialize.SweepTask()
	core.RunWindowsServer()
}
