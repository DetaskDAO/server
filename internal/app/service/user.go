package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/request"
	"code-market-admin/internal/app/model/response"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUserInfo
// @function: CreateUserInfo
// @description: 创建用户信息
// @param: createuserInfo request.CreateUserInfoRequest
// @return: err error
func CreateUserInfo(c *gin.Context, userInfo request.CreateUserInfoRequest) (err error) {
	chain := GetChain(c) // 获取链
	if err = global.MAPDB[chain].Model(&model.User{}).Save(&userInfo.User).Error; err != nil {
		return err
	}
	return nil
}

// GetUserAvatar
// @function: GetUserAvatar
// @description: 获取个⼈资料(用户名和头像)
// @param: userAvatar request.GetUserInfoRequest
// @return: err error, user model.User
func GetUserAvatar(c *gin.Context, userAvatar request.GetUserInfoRequest) (err error, user model.User) {
	chain := GetChain(c) // 获取链
	db := global.MAPDB[chain].Model(&model.User{})
	if err = db.Where("address = ?", userAvatar.Address).Find(&user).Error; err != nil {
		return err, user
	}
	return err, user
}

// GetUserInfo
// @function: GetUserInfo
// @description: 获取个⼈资料
// @param: userInfo request.GetUserInfoRequest
// @return: err error, user model.User
func GetUserInfo(c *gin.Context, userInfo request.GetUserInfoRequest) (err error, user model.User) {
	chain := GetChain(c) // 获取链
	if err = global.MAPDB[chain].Model(&model.User{}).Where("address = ?", userInfo.Address).Find(&user).Error; err != nil {
		return err, user
	}
	return err, user
}

// UpdateUserInfo
// @function: UpdateUserInfo
// @description: 修改个⼈资料
// @param: updateuserInfo request.UpdateUserInfoRequest
// @return: err error
func UpdateUserInfo(c *gin.Context, userInfo request.UpdateUserInfoRequest) (err error) {
	chain := GetChain(c)             // 获取链
	userInfo.Address = GetAddress(c) //
	if err = global.MAPDB[chain].Model(&model.User{}).Where("address = ?", userInfo.Address).Updates(&userInfo.User).Error; err != nil {
		return err
	}
	return nil
}

// UnReadMsgCount
// @description: 获取未读消息数量
// @param: userID string
// @return: count int64, err error
func UnReadMsgCount(c *gin.Context, userID uint) (count int64, err error) {
	chain := GetChain(c) // 获取链
	err = global.MAPDB[chain].Model(&model.Message{}).Where("status = 0 AND type = 0 AND rec_id = ?", userID).Count(&count).Error
	return count, err
}

// UnReadMsgList
// @description: 获取未读消息
// @param: userID string
// @return: count int64, err error
func UnReadMsgList(c *gin.Context, userID uint, searchInfo request.UnReadRequest) (list []response.MsgListRespond, total int64, err error) {
	chain := GetChain(c) // 获取链
	lang := GetLang(c)   // 获取语言
	db := global.MAPDB[chain].Model(&model.Message{})
	if lang == "zh" {
		db.Select("*,message_zh as message")
	}
	db = db.Where("status = 0 AND rec_id = ?", userID)
	if searchInfo.OrderId != 0 && searchInfo.Type != 0 {
		db = db.Where("order_id = ? AND type = ?", searchInfo.OrderId, searchInfo.Type)
	} else {
		db = db.Where("type = 0")
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Preload("User").Order("created_at desc").Limit(10).Find(&list).Error
	return list, total, err
}

// ReadMsg
// @description: 阅读信息
// @param: userID string, msgID uint
// @return: err error
func ReadMsg(c *gin.Context, userID uint, readMsgReq request.ReadMsgRequest) (err error) {
	chain := GetChain(c) // 获取链
	if readMsgReq.OrderId == 0 {
		raw := global.MAPDB[chain].Model(&model.Message{}).Where("rec_id = ? AND id = ?", userID, readMsgReq.ID).Update("status", 1)
		if raw.RowsAffected == 0 || raw.Error != nil {
			return errors.New("修改状态失败")
		}
	} else {
		raw := global.MAPDB[chain].Model(&model.Message{}).Where("rec_id = ? AND order_id = ?", userID, readMsgReq.OrderId).Update("status", 1)
		if raw.RowsAffected == 0 || raw.Error != nil {
			return errors.New("修改状态失败")
		}
	}

	return nil
}

// ReadAllMsg
// @description: 全部已读
// @param: userID string
// @return: err error
func ReadAllMsg(c *gin.Context, userID uint) (err error) {
	chain := GetChain(c) // 获取链
	raw := global.MAPDB[chain].Model(&model.Message{}).Where("rec_id = ?", userID).Where("type = 0").Update("status", 1)
	if raw.RowsAffected == 0 {
		return errors.New("修改状态失败")
	}
	return raw.Error
}

// MsgList
// @description: 分页获取消息
// @param: userID string, msgID uint
// @return: err error
func MsgList(c *gin.Context, searchInfo request.MsgListRequest, userID uint) (list []response.MsgListRespond, total int64, err error) {
	chain := GetChain(c) // 获取链
	lang := GetLang(c)   // 获取语言
	db := global.MAPDB[chain].Model(&model.Message{}).Where("type = 0")
	if lang == "zh" {
		db.Select("*,message_zh as message")
	}
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	db = db.Where("rec_id = ?", userID)

	if err = db.Count(&total).Error; err != nil {
		return list, total, err
	}
	db = db.Limit(limit).Offset(offset)
	err = db.Preload("User").Order("created_at desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return list, total, err
	}
	return list, total, nil
}
