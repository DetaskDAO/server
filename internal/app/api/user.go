package api

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model/request"
	"code-market-admin/internal/app/model/response"
	"code-market-admin/internal/app/service"
	"code-market-admin/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateUserInfo
// @Tags UserApi
// @Summary 创建个人资料
// @accept application/json
// @Produce application/json
// @Router /user/createUserInfo [get]
func CreateUserInfo(c *gin.Context) {
	var userInfo request.CreateUserInfoRequest
	_ = c.ShouldBindJSON(&userInfo)
	// 检验字段
	if err := utils.Verify(userInfo.User, utils.CreateUserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateUserInfo(c, userInfo); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// GetUserAvatar
// @Tags UserApi
// @Summary 获取个人资料(用户名和头像)
// @accept application/json
// @Produce application/json
// @Router /user/getUserAvatar [get]
func GetUserAvatar(c *gin.Context) {
	var userAvatar request.GetUserInfoRequest
	_ = c.ShouldBindQuery(&userAvatar)
	// 校验字段
	if err := utils.Verify(userAvatar.User, utils.UserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, userRes := service.GetUserAvatar(c, userAvatar); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else if userRes.Address == "" {
		response.FailWithMessage("查无此人", c)
	} else {
		response.OkWithDetailed(response.UserAvatarRespond{
			Username: *userRes.Username,
			Avatar:   *userRes.Avatar,
		}, "获取成功", c)
	}
}

// GetUserInfo
// @Tags UserApi
// @Summary 获取个人资料
// @accept application/json
// @Produce application/json
// @Router /user/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
	var userInfo request.GetUserInfoRequest
	_ = c.ShouldBindQuery(&userInfo)
	// 校验字段
	if err := utils.Verify(userInfo.User, utils.UserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, userRes := service.GetUserInfo(c, userInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else if userRes.Address == "" {
		response.FailWithMessage("查无此人", c)
	} else {
		response.OkWithDetailed(userRes, "获取成功", c)
	}
}

// UpdateUserInfo
// @Tags UserApi
// @Summary 修改个人资料
// @accept application/json
// @Produce application/json
// @Router /user/updateUserInfo [get]
func UpdateUserInfo(c *gin.Context) {
	var userInfo request.UpdateUserInfoRequest
	_ = c.ShouldBindJSON(&userInfo)
	if err := service.UpdateUserInfo(c, userInfo); err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// UnReadMsgCount
// @Tags UserApi
// @Summary 获取未读消息数量
// @Router /user/unReadMsgCount [get]
func UnReadMsgCount(c *gin.Context) {
	userID := c.GetUint("userID") // 操作人ID
	if count, err := service.UnReadMsgCount(c, userID); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]int64{"count": count}, "获取成功", c)
	}
}

// UnReadMsgList
// @Tags UserApi
// @Summary 获取未读消息
// @Router /user/unReadMsgList [get]
func UnReadMsgList(c *gin.Context) {
	var searchInfo request.UnReadRequest
	_ = c.ShouldBindQuery(&searchInfo)
	userID := c.GetUint("userID") // 操作人ID
	if list, total, err := service.UnReadMsgList(c, userID, searchInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// ReadMsg
// @Tags UserApi
// @Summary 阅读信息
// @Router /user/readMsg [post]
func ReadMsg(c *gin.Context) {
	var readMsgReq request.ReadMsgRequest
	_ = c.ShouldBindJSON(&readMsgReq)
	userID := c.GetUint("userID") // 操作人ID
	if err := service.ReadMsg(c, userID, readMsgReq); err != nil {
		global.LOG.Error("操作失败!", zap.Error(err))
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}

// ReadAllMsg
// @Tags UserApi
// @Summary 阅读信息
// @Router /user/readAllMsg [post]
func ReadAllMsg(c *gin.Context) {
	var IDReq request.IDRequest
	_ = c.ShouldBindJSON(&IDReq)
	userID := c.GetUint("userID") // 操作人ID
	if err := service.ReadAllMsg(c, userID); err != nil {
		global.LOG.Error("操作失败!", zap.Error(err))
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}

// MsgList
// @Tags UserApi
// @Summary 分页获取消息
// @Router /user/msgList [get]
func MsgList(c *gin.Context) {
	var searchInfo request.MsgListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	userID := c.GetUint("userID") // 操作人ID
	if list, total, err := service.MsgList(c, searchInfo, userID); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, "获取成功", c)
	}
}
