package model

import (
	"code-market-admin/internal/app/global"
	"time"
)

type Apply struct {
	global.MODEL
	ApplyAddr string    `gorm:"column:apply_addr;index:,unique,composite:apply_addr_task_id;type:char(42)" json:"apply_addr" form:"apply_addr"` // 申请人地址
	TaskID    uint64    `gorm:"column:task_id;index:,unique,composite:apply_addr_task_id" json:"task_id" form:"task_id"`                        // 任务ID
	OrderId   int64     `gorm:"column:order_id" json:"order_id" form:"order_id"`
	Price     string    `gorm:"column:price" json:"price" form:"price"`              // 报价
	Desc      string    `gorm:"column:desc" json:"desc" form:"desc"`                 // 自我介绍
	SortTime  time.Time `gorm:"column:sort_time" json:"sort_time" form:"sort_time"`  // 不感兴趣
	Status    uint8     `gorm:"column:status;default:1" json:"status" form:"status"` // 状态 1: 报名中 2:甲方已选择
}
