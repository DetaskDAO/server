package model

import (
	"code-market-admin/internal/app/global"
	"github.com/lib/pq"
)

type User struct {
	global.MODEL
	Address     string        `gorm:"column:address;type:char(42);UNIQUE;comment:用户地址" json:"address" form:"address"`
	Username    *string       `gorm:"column:username;type:varchar(42);UNIQUE;comment:用户名" json:"username" form:"username"`
	Description *string       `gorm:"column:description;type:varchar(100);comment:自我介绍" json:"description" form:"description"`
	Avatar      *string       `gorm:"column:avatar;type:varchar(200);comment:用户头像" json:"avatar" form:"avatar"`
	Telegram    *string       `gorm:"column:telegram;type:varchar(100);comment:Telegram" json:"telegram" form:"telegram"`
	Wechat      *string       `gorm:"column:wechat;type:varchar(100);comment:Wechat" json:"wechat" form:"wechat"`
	Skype       *string       `gorm:"column:skype;type:varchar(100);comment:Skype" json:"skype" form:"skype"`
	Discord     *string       `gorm:"column:discord;type:varchar(100);comment:Discord" json:"discord" form:"discord"`
	Phone       *string       `gorm:"column:phone;type:varchar(100);comment:Phone" json:"phone" form:"phone"`
	Role        pq.Int64Array `gorm:"column:role;type:integer[];default:'{}';comment:技能" json:"role" form:"role"` // 技能
}
