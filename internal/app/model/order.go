package model

import (
	"code-market-admin/internal/app/global"
	"time"
)

type Order struct {
	global.MODEL
	OrderId     int64     `gorm:"column:order_id;unique;comment:操作者" json:"order_id" form:"order_id"`
	TaskID      int64     `gorm:"column:task_id;comment:操作者" json:"task_id" form:"task_id"`
	StartAt     time.Time `gorm:"column:start_at;comment:开始时间" json:"start_at"` // 开始时间
	Task        string    `gorm:"column:task;comment:需求详情" json:"task" form:"task"`
	Issuer      string    `gorm:"column:issuer;type:char(42)" json:"issuer" form:"issuer"`
	Worker      string    `gorm:"column:worker;type:char(42)" json:"worker" form:"worker"`
	Attachment  string    `gorm:"column:attachment;default:'';size:150" json:"attachment" form:"attachment"`
	Signature   string    `gorm:"column:signature;type:varchar(132)" json:"signature" form:"signature"`
	SignAddress string    `gorm:"column:sign_address;type:varchar(42)" json:"sign_address" form:"sign_address"`
	SignNonce   int64     `gorm:"column:sign_nonce" json:"sign_nonce" form:"sign_nonce"`
	Currency    string    `gorm:"column:currency;type:char(42);comment:币种" json:"currency" form:"currency"`                 // 币种
	Amount      string    `gorm:"column:amount;comment:金额" json:"amount" form:"amount"`                                     // 金额
	Pending     uint      `gorm:"column:pending;comment:待领取余额" json:"pending" form:"pending"`                               // 待领取余额
	Stages      string    `gorm:"column:stages;comment:阶段JSON" json:"stages" form:"stages"`                                 // 阶段JSON
	PayType     uint8     `gorm:"column:pay_type;comment:付款方式 0:Unknown 1:Due 2:Confirm" json:"pay_type" form:"pay_type"`   // 付款方式 0: Unknown 1: Due 2: Confirm
	State       *uint8    `gorm:"column:state;default:0;comment:任务状态 0:进行中 1:已完成" json:"state" form:"state"`                // 任务状态 0:进行中 1: 已完成
	Progress    uint8     `gorm:"column:progress;comment:阶段" json:"progress" form:"progress"`                               // 阶段
	Status      string    `gorm:"column:status;default:'WaitWorkerStage';size:30;comment:事件状态" json:"status" form:"status"` // 事件状态 10：
}
