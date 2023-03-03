package global

import (
	"code-market-admin/internal/app/config"
	"github.com/allegro/bigcache/v3"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	StartTime     time.Time                 // 记录运行时间
	DB            *gorm.DB                  // 数据库链接
	MAPDB         map[string]*gorm.DB       // 多链数据库链接
	ProviderMap   map[string]string         // 多链Provider RPC
	ProviderIPFS  string                    // IPFS 节点
	CONFIG        config.Server             // 配置信息
	LOG           *zap.Logger               // 日志框架
	TokenCache    *bigcache.BigCache        // Token 缓存
	JsonCache     *bigcache.BigCache        // JSON 缓存
	ContractAddr  map[string]common.Address // 合约地址
	ContractEvent map[common.Hash]string    // 合约事件
)
