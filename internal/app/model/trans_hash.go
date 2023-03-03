package model

import (
	"code-market-admin/internal/app/global"
)

type TransHash struct {
	global.MODEL
	Hash      string `gorm:"column:hash;type:char(68);unique;not null;comment:交易hash" json:"hash" form:"hash"` // 交易hash唯一
	EventName string `gorm:"column:event_name;size:20;comment:事件名称" json:"event_name"`
	SendAddr  string `gorm:"column:send_addr;type:char(42);comment:提交人" json:"send_addr" form:"send_addr"`
	Raw       string `gorm:"column:raw;comment:原始JSON" json:"raw" form:"raw"`
	Status    uint8  `gorm:"column:status;default:0;comment:0:处理中 1:交易成功 2:交易失败 3:超过解析次数 4:交易成功未匹配事件" json:"status" form:"status"` // 状态 0 处理中 1 交易成功 2 交易失败 3 超过解析次数 4 交易成功未匹配事件
}
