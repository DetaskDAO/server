package model

import "code-market-admin/internal/app/global"

type Message struct {
	global.MODEL
	OrderId   int64  `gorm:"column:order_id" json:"order_id" form:"order_id"`       // Order ID
	SendID    uint   `json:"send_id" form:"send_id"`                                // 发送者ID
	RecID     uint   `json:"rec_id" form:"rec_id"`                                  // 接收者ID
	Message   string `gorm:"column:message;size:500" json:"message" form:"message"` // 站内信内容
	MessageZh string `gorm:"column:message_zh;size:500" json:"-" form:"-"`          // 中文站内信内容
	Status    uint8  `gorm:"column:status" json:"status" form:"status"`             // 站内信的查看状态 0: 未读 1已读
	Type      uint8  `gorm:"column:type;default:0" json:"type" form:"type"`         // 消息类型 0: 正常消息 1: 提示消息
}
