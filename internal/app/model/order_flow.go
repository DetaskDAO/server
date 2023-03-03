package model

import (
	"time"
)

type OrderFlow struct {
	ID          uint      `gorm:"primarykey"`                                                            // 主键ID
	CreatedAt   time.Time `json:"created_at"`                                                            // 创建时间
	Del         uint8     `gorm:"column:del;default:0;comment:删除状态 0: 未删除 1: 已删除" json:"del" form:"del"` // 删除状态 0: 未删除 1: 已删除
	OrderId     int64     `gorm:"column:order_id;index:,unique,composite:order_id_level;comment:删除" json:"order_id" form:"order_id"`
	Level       int64     `gorm:"column:level;default:1;index:,unique,composite:order_id_level;comment:节点" json:"level" form:"level"` // 节点
	Operator    string    `gorm:"column:operator;type:char(42);comment:操作者" json:"operator" form:"operator"`                          // 操作者
	Signature   string    `gorm:"column:signature;default:'';type:varchar(132);comment:签名" json:"signature" form:"signature"`
	SignAddress string    `gorm:"column:sign_address;type:varchar(42);comment:签名地址" json:"sign_address" form:"sign_address"`
	SignNonce   int64     `gorm:"column:sign_nonce;comment:签名 Nonce" json:"sign_nonce" form:"sign_nonce"`
	Obj         string    `gorm:"column:obj" json:"obj" form:"obj"`
	Attachment  string    `gorm:"column:attachment" json:"attachment" form:"attachment"`
	Stages      string    `gorm:"column:stages" json:"stages" form:"stages"`
	Status      string    `gorm:"column:status;default:'WaitWorkerStage';size:30;comment:事件状态" json:"status" form:"status"` // 事件状态
	Audit       uint8     `gorm:"column:audit;default:0;comment:审核状态 0:未审核 1:已审核" json:"audit" form:"audit"`                // 审核状态 0: 未审核 1: 已审核
}
