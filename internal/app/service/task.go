package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/request"
	"code-market-admin/internal/app/model/response"
	_ "code-market-admin/internal/app/model/response"
	"code-market-admin/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

// GetTaskList
// @function: GetTaskList
// @description: 分页获取需求数据
// @param: task model.Tasks, info Req.PageInfo
// @return: err error, list interface{}, total int64
func GetTaskList(c *gin.Context, searchInfo request.GetTaskListRequest) (err error, list interface{}, total int64) {
	chain := GetChain(c) // 获取链
	var responses []response.GetTaskListRespond
	var taskList []model.Task
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	db := global.MAPDB[chain].Model(&model.Task{})
	// 根据ID过滤
	if searchInfo.ID != 0 {
		db = db.Where("id = ?", searchInfo.ID)
	}
	// 根据链上TaskID过滤
	if searchInfo.TaskID != 0 {
		db = db.Where("task_id = ?", searchInfo.TaskID)
	}
	// 根据标题过滤
	if searchInfo.Title != "" {
		db = db.Where("title ILIKE ?", "%"+searchInfo.Title+"%")
	}
	// 根据创建者Hash过滤
	if searchInfo.Issuer != "" {
		db = db.Where("issuer = ?", searchInfo.Issuer)
	} else {
		//报名开关: 0.关  1.开  这里前端显示
		db = db.Where("apply_switch = 1")
	}
	// 根据技能要求过滤
	if searchInfo.Role != nil {
		db = db.Where("role && ?", searchInfo.Role)
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, list, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("created_at desc").Find(&taskList).Error
	}
	if searchInfo.TaskID == 0 {
		return err, taskList, total
	}
	// 获取 IPFS
	for _, task := range taskList {
		res := response.GetTaskListRespond{Task: task}
		// 获取IPFS
		hash := res.Attachment
		res.Attachment, _ = utils.GetJSONFromCid(hash)
		responses = append(responses, res)
	}
	return err, responses, total
}

// CreateTask
// @function: CreateTask
// @description: 发布需求
// @param: taskReq request.CreateTaskRequest
// @return: err error
func CreateTask(c *gin.Context, taskReq request.CreateTaskRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	if err != nil {
		return errors.New("新建失败")
	}
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "TaskCreated", Hash: taskReq.Hash}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

// UpdatedTask
// @function: UpdatedTask
// @description: 修改需求
// @param: task model.Tasks, info Req.PageInfo
// @return: err error, list interface{}, total int64
func UpdatedTask(c *gin.Context, taskReq request.UpdatedTaskRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	if err != nil {
		return errors.New("新建失败")
	}
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "TaskModified", Hash: taskReq.Hash}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

// ModifyApplySwitch 修改报名开关
func ModifyApplySwitch(c *gin.Context, req request.ModifyApplySwitchRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	res := global.MAPDB[chain].Model(&model.Task{}).Where("task_id = ? AND issuer = ?", req.TaskID, address).Update("apply_switch", req.ApplySwitch)
	if res.RowsAffected == 0 {
		return errors.New("操作失败")
	}
	return res.Error
}

// GetSillTreeMap
// @description: 获取技能树
// @param: sillId uint
// @return: res []*model.Skill, err error
func GetSillTreeMap(sillId uint) (res []*model.Skill, err error) {
	var allMenus []*model.Skill
	err = global.DB.Order("sort").Find(&allMenus).Error
	if err != nil {
		return
	}
	return tree(allMenus, sillId), err
}

// tree
// @description: 生成技能树结构
// @param: sills []*model.Skill, pid uint
// @return: []*model.Skill
func tree(sills []*model.Skill, pid uint) []*model.Skill {
	// 节点
	var nodes []*model.Skill
	if reflect.ValueOf(sills).IsValid() {
		// 循环所有 一级技能
		for _, v := range sills {
			if v.ParentId == pid {
				// 递归方式循环子树
				v.Children = append(v.Children, tree(sills, v.ID)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}
