package model

import (
	"code-market-admin/internal/app/global"
	"github.com/lib/pq"
)

type Task struct {
	global.MODEL
	TaskID      uint64        `gorm:"column:task_id;unique;comment:链上任务ID" json:"task_id" form:"task_id"`                          // 链上任务ID，唯一值
	Issuer      string        `gorm:"column:issuer;type:char(42);comment:发布人hash" json:"issuer" form:"issuer"`                     // 发布人hash
	Title       string        `gorm:"column:title;comment:标题" json:"title" form:"title"`                                           // 标题
	Period      uint32        `gorm:"column:period;comment:预计周期" json:"period" form:"period"`                                      // 预计周期
	Budget      string        `gorm:"column:budget;comment:预计金额" json:"budget" form:"budget"`                                      // 预计金额
	Role        pq.Int64Array `gorm:"column:role;type:integer[];comment:所需技能" json:"role" form:"role"`                             // 所需技能
	ApplyCount  uint          `gorm:"column:apply_count;default:0;comment:报名人数" json:"apply_count" form:"apply_count"`             // 报名人数
	Attachment  string        `gorm:"column:attachment;comment:描述IPFS" json:"attachment" form:"attachment"`                        // 描述IPFS
	Currency    string        `gorm:"column:currency;size:30;comment:币种" json:"currency" form:"currency"`                          // 币种
	ApplySwitch uint8         `gorm:"column:apply_switch;default:1;comment:报名开关: 0.关 1.开" json:"apply_switch" form:"apply_switch"` // 报名开关: 0.关  1.开
}
