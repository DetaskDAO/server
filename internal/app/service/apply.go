package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/request"
	"code-market-admin/internal/app/model/response"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

// GetApplyList
// @function: GetApplyList
// @description: 获取需求里报名详情
// @param: c *gin.Context, searchInfo request.GetApplyListRequest
// @return: err error, list interface{}, total int64
func GetApplyList(c *gin.Context, searchInfo request.GetApplyListRequest) (err error, list interface{}, total int64) {
	chain := GetChain(c) // 获取链
	var applyList []response.GetApplyListRespond
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	db := global.MAPDB[chain].Model(&model.Apply{})
	// 根据TaskId过滤
	if searchInfo.TaskID != 0 {
		db = db.Where("task_id = ?", searchInfo.TaskID).Order("sort_time asc")
	}
	if searchInfo.Status != 0 {
		db = db.Where("status = ?", searchInfo.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, list, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("created_at desc").Preload("User").Find(&applyList).Error
	}
	return err, applyList, total
}

// GetApply
// @function: GetApply
// @description: 分页获取个⼈报名中的项⽬
// @param: c *gin.Context, searchInfo request.GetApplyRequest
// @return: err error, list interface{}, total int64
func GetApply(c *gin.Context, searchInfo request.GetApplyRequest) (err error, list interface{}, total int64) {
	chain := GetChain(c) // 获取链
	var applyList []response.GetApplyRespond
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	db := global.MAPDB[chain].Model(&model.Apply{})
	// 根据报名人地址过滤
	if searchInfo.TaskID != 0 {
		db = db.Where("task_id = ?", searchInfo.TaskID)
	}
	if searchInfo.ApplyAddr != "" {
		db = db.Where("apply_addr = ?", searchInfo.ApplyAddr)
	}
	if searchInfo.Status != 0 {
		db = db.Where("status = ?", searchInfo.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, list, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("created_at desc").Preload("Task").Find(&applyList).Error
	}
	return err, applyList, total
}

// CreateApply
// @function: CreateApply
// @description: 添加报名信息
// @param: c *gin.Context, applyReq request.CreateApplyRequest
// @return: err error
func CreateApply(c *gin.Context, applyReq request.CreateApplyRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	// 保存请求数据
	raw, err := json.Marshal(applyReq)
	if err != nil {
		return errors.New("新建失败")
	}
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "ApplyFor", Hash: applyReq.Hash, Raw: string(raw)}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

// UpdatedApply
// @function: UpdatedApply
// @description: 修改报名信息
// @param: applyReq request.UpdatedApplyRequest
// @return: err error
func UpdatedApply(c *gin.Context, applyReq request.UpdatedApplyRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	// 保存请求数据
	raw, err := json.Marshal(applyReq)
	if err != nil {
		return errors.New("新建失败")
	}
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "ApplyFor", Hash: applyReq.Hash, Raw: string(raw)}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

// DeleteApply
// @function: DeleteApply
// @description: 删除报名信息
// @param: applyReq request.DeleteApplyRequest
// @return: err error
func DeleteApply(c *gin.Context, hash string) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "CancelApply", Hash: hash}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

// UpdatedApplySort
// @function: UpdatedApplySort
// @description: 更新报名列表排序
// @param: taskReq request.CreateTaskRequest
// @return: err error
func UpdatedApplySort(c *gin.Context, sortReq request.UpdatedApplySortRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	db := global.MAPDB[chain].Model(&model.Apply{})
	// 是否任务本人操作
	if err = global.MAPDB[chain].Model(&model.Task{}).Where("task_id = ? AND issuer = ?", sortReq.TaskID, address).First(&model.Task{}).Error; err != nil {
		return errors.New("设置失败")
	}
	// 更新状态
	raw := db.Where("task_id = ?", sortReq.TaskID).Where("apply_addr = ?", sortReq.ApplyAddr).Update("sort_time", time.Now())
	if raw.RowsAffected == 0 {
		return errors.New("设置失败")
	}
	return raw.Error
}
