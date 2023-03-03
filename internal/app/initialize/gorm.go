package initialize

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// InitChainDB 多链数据库
func InitChainDB() {
	global.MAPDB = make(map[string]*gorm.DB) // 初始化map
	for _, chain := range global.CONFIG.BlockChain {
		db := GormPgSql(chain.Name)
		if db != nil {
			global.MAPDB[chain.Name] = db
			RegisterTables(global.MAPDB[chain.Name]) // 初始化表
		}
	}
}

// InitCommonDB 通用数据库
func InitCommonDB() {
	db := GormPgSql("")
	if db != nil {
		global.DB = db
		RegisterCommonTables(db) // 初始化表
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Apply{},
		model.Order{},
		model.User{},
		model.TransHash{},
		model.Task{},
		model.Upload{},
		model.Message{},
		model.OrderFlow{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}

// RegisterCommonTables 注册数据库表专用
func RegisterCommonTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Skill{},
		model.MessageTmpl{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
