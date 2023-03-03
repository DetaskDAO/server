package response

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
)

type UserAvatarRespond struct {
	Address  string `gorm:"column:address" json:"address" form:"address"`
	Username string `gorm:"column:username" json:"username" form:"username"`
	Avatar   string `gorm:"column:avatar" json:"avatar" form:"avatar"`
}

type user struct {
	global.MODEL
	Username string `gorm:"column:username" json:"username" form:"username"`
	Avatar   string `gorm:"column:avatar" json:"avatar" form:"avatar"`
	Address  string `gorm:"column:address;type:char(42);UNIQUE" json:"address" form:"address"`
}

type MsgListRespond struct {
	model.Message
	User user `json:"user" form:"user" gorm:"foreignKey:ID;references:SendID"`
}
