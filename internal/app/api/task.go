package api

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/request"
	response "code-market-admin/internal/app/model/response"
	"code-market-admin/internal/app/service"
	"code-market-admin/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetTaskList
// @Tags TaskApi
// @Summary 分页获取需求数据
// @accept application/json
// @Produce application/json
// @Router /task/getTaskList [get]
func GetTaskList(c *gin.Context) {
	var searchInfo request.GetTaskListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	// 校验字段
	if err := utils.Verify(searchInfo.PageInfo, utils.PageSizeLimitVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetTaskList(c, searchInfo); err != nil {
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

// CreateTask
// @Tags TaskApi
// @Summary 发布需求
// @accept application/json
// @Produce application/json
// @Router /task/createTask [post]
func CreateTask(c *gin.Context) {
	var task request.CreateTaskRequest
	_ = c.ShouldBindJSON(&task)
	// 校验字段
	if err := utils.Verify(task.Task, utils.CreateTaskVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := service.CreateTask(c, task); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// UpdatedTask
// @Tags TaskApi
// @Summary 更新需求信息
// @accept application/json
// @Produce application/json
// @Router /task/UpdatedTask [post]
func UpdatedTask(c *gin.Context) {
	var task request.UpdatedTaskRequest
	_ = c.ShouldBindJSON(&task)

	if err := utils.Verify(task.Task, utils.UpdatedTaskVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdatedTask(c, task); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ModifyApplySwitch
// @Tags TaskApi
// @Summary 修改报名开关
// @accept application/json
// @Produce application/json
// @Router /task/modifyApplySwitch [post]
func ModifyApplySwitch(c *gin.Context) {
	var task request.ModifyApplySwitchRequest
	_ = c.ShouldBindJSON(&task)
	if err := utils.Verify(task, utils.ModifyApplySwitchVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ModifyApplySwitch(c, task); err != nil {
		global.LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// GetSillTreeMap
// @Tags TaskApi
// @Summary 获取技能树
// @accept application/json
// @Produce application/json
// @Router /task/getSillTreeMap [get]
func GetSillTreeMap(c *gin.Context) {
	var searchInfo model.Skill
	_ = c.ShouldBindQuery(&searchInfo)
	if list, err := service.GetSillTreeMap(searchInfo.ID); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}
