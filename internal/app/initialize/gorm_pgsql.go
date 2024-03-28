package initialize

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/initialize/internal"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql(Prefix string) *gorm.DB {
	p := global.CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(Prefix)); err != nil {
		global.LOG.Error("Postgres connect error", zap.String("err", err.Error()))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}
