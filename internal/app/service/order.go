package service

import (
	"code-market-admin/internal/app/blockchain"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/message"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/model/request"
	"code-market-admin/internal/app/model/response"
	"code-market-admin/internal/app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// GetOrderList
// @function: GetOrderList
// @description: 分页获取个⼈进⾏中的项⽬
// @param: searchInfo request.GetOrderListRequest
// @return: err error, list interface{}, total int64
func GetOrderList(c *gin.Context, searchInfo request.GetOrderListRequest) (err error, list interface{}, total int64) {
	chain := GetChain(c) // 获取链
	var responses []response.GetOrderListResponse
	var orderList []response.GetOrderListResponse
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)
	db := global.MAPDB[chain].Model(&model.Order{}).Where(searchInfo.Order)
	// 获取数量
	err = db.Count(&total).Error
	if err != nil {
		return err, list, total
	}
	// 分页获取
	db = db.Limit(limit).Offset(offset)
	err = db.Order("created_at desc").Preload("IssuerInfo").Preload("WorkerInfo").Find(&orderList).Error
	// 根据订单ID过滤 需要获取IPFS 具体内容
	if searchInfo.OrderId != 0 && len(orderList) == 1 && orderList[0].Attachment != "" {
		// 从缓存获取
		orderList[0].StageJson, _ = utils.GetJSONFromCid(orderList[0].Attachment)
		// WaitProlongAgree || WaitAppendAgree 状态需要 返回原始数据（上次数据）
		if orderList[0].Status == "WaitProlongAgree" || orderList[0].Status == "WaitAppendAgree" {
			global.MAPDB[chain].Model(&model.OrderFlow{}).Select("stages").Where("order_id = ? AND status = 'IssuerAgreeStage' AND del = 0", orderList[0].OrderId).Order("level desc").First(&orderList[0].LastStages)
			var attachment string
			global.MAPDB[chain].Model(&model.OrderFlow{}).Select("attachment").Where("order_id = ? AND status = 'IssuerAgreeStage' AND del = 0", orderList[0].OrderId).Order("level desc").First(&attachment)
			// 从缓存获取
			orderList[0].LastStageJson, _ = utils.GetJSONFromCid(attachment)
		}
	}
	// IPFS
	for _, item := range orderList {
		if errUn := json.Unmarshal([]byte(item.Order.Task), &item.Task); errUn != nil {
			fmt.Println(errUn)
			continue
		}
		// 获取IPFS
		// 从缓存获取
		hash := item.Task.Attachment
		item.Task.Attachment, _ = utils.GetJSONFromCid(hash)
		responses = append(responses, item)
	}
	return err, responses, total
}

// CreateOrder
// @function: CreateOrder
// @description: 添加任务
// @param: taskReq request.CreateTaskRequest
// @return: err error
func CreateOrder(c *gin.Context, orderReq request.CreateOrderRequest, address string) (err error) {
	chain := GetChain(c) // 获取链
	// 保存交易hash
	transHash := model.TransHash{SendAddr: address, EventName: "OrderCreated", Hash: orderReq.Hash}
	if err = SaveHash(transHash, chain); err != nil {
		return errors.New("新建失败")
	}
	return nil
}

func DeleteOrder(c *gin.Context, orderReq request.DeleteOrderRequest) (err error) {
	chain := GetChain(c)     // 获取链
	address := GetAddress(c) // 操作人
	var order model.Order
	tx := global.MAPDB[chain].Begin()
	// 删除order
	raw := tx.Model(&model.Order{}).Where("order_id = ? AND (issuer = ? OR worker = ?) AND progress < 2", orderReq.OrderId, address, address).Delete(&order)
	if raw.RowsAffected == 0 || raw.Error != nil {
		tx.Rollback()
		return errors.New("删除失败")
	}
	// 更新status
	raw = tx.Model(&model.Apply{}).Where("order_id = ?", orderReq.OrderId).Update("status", 1)
	if raw.RowsAffected != 1 || raw.Error != nil {
		tx.Rollback()
		return errors.New("删除失败")
	}

	return tx.Commit().Error
}

// UpdatedStage
// @function: CreatedStage
// @description: 更新阶段划分
// @param: stage request.UpdatedStageRequest
// @return: err error
func UpdatedStage(c *gin.Context, stage request.UpdatedStageRequest, address string) (err error) {
	chain := GetChain(c) // 获取链
	// TODO: 权限控制
	// 保存交易Hash
	runBool, err := storeHash(stage, address, chain)
	if err != nil || runBool {
		return err
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	fmt.Println("开始事务")
	// 查询当前记录
	var order model.Order
	if err = tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).First(&order).Error; err != nil {
		return err
	}
	// 查询日志记录
	var orderFlowTop model.OrderFlow
	if err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND del = 0", stage.OrderId).Order("level desc").First(&orderFlowTop).Error; err != nil {
		return err
	}
	// 状态流转 校验正确性
	if stage.Status != "WaitWorkerStage" && order.Status != stage.Status && stage.Status != "" {
		if err = statusValid(order.Status, stage.Status, stage, order); err != nil {
			tx.Rollback()
			return err
		}
	}
	// 消息通知
	if err = sendMessage(tx, order, stage, orderFlowTop, address, chain); err != nil {
		tx.Rollback()
		return err
	}
	// rollback 操作
	err, okRun := rollbackStatus(tx, order.Status, stage.Status, stage)
	if err != nil {
		tx.Rollback()
		return err
	}
	if okRun {
		return tx.Commit().Error
	}
	// 需要更新Obj
	if stage.Obj != "" && gjson.Get(orderFlowTop.Obj, "stages").String() != gjson.Get(stage.Obj, "stages").String() {
		// 查询当前attachment
		var attachment string
		err = global.MAPDB[chain].Model(&model.Order{}).Select("attachment").Where("order_id = ?", stage.OrderId).First(&attachment).Error
		if err != nil {
			return err
		}
		// 写入last字段
		stage.Obj, _ = sjson.Set(stage.Obj, "last", attachment)
		// 上传JSON获取IPFS CID
		err, hashJSON := IPFSUploadJSON(stage.Obj)
		if err != nil {
			return err
		}
		// 更新数据
		stage.Attachment = hashJSON
	}
	// 更新Order表
	raw := tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).Updates(&stage.Order)
	if raw.RowsAffected == 0 || raw.Error != nil {
		tx.Rollback()
		return errors.New("创建失败")
	}
	// 保存Order流程记录
	if err = saveFlow(tx, stage, address); err != nil {
		tx.Rollback()
		return err
	}
	// 处理特殊状态(WaitAppendAgree状态提交交付物)
	if err = dealWaitAppendAgree(tx, order, stage, orderFlowTop, address); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdatedProgress
// @description: 更新阶段状态
// @param: updatedProgress request.UpdatedProgressRequest
// @return:  err error
func UpdatedProgress(c *gin.Context, updatedProgress request.UpdatedProgressRequest) (err error) {
	chain := GetChain(c) // 获取链
	for i := 0; i < 3; i++ {
		operation := func() error {
			return blockchain.UpdatedProgress(updatedProgress.OrderId, chain)
		}
		if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)); err != nil {
			global.LOG.Error("UpdatedProgress service error", zap.Error(err))
			continue
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}

// WaitAppendAgree 状态提交交付处理 同时只会提交一个交付物
func dealWaitAppendAgree(tx *gorm.DB, order model.Order, stage request.UpdatedStageRequest, orderFlowTop model.OrderFlow, address string) (err error) {
	if stage.Status != "" || order.Status != "WaitAppendAgree" {
		return
	}
	// 判断是否有新交付物
	attachmentNew := gjson.Get(stage.Obj, "stages")
	attachmentOld := gjson.Get(orderFlowTop.Obj, "stages")
	// 不是阶段交付
	if len(attachmentNew.Array()) != len(attachmentOld.Array()) {
		return
	}
	// 遍历得到哪个阶段交付
	var index int
	for i := 0; i < len(attachmentNew.Array()); i++ {
		if attachmentNew.Array()[i].String() != attachmentOld.Array()[i].String() && attachmentNew.Array()[i].String() != "" {
			index = i
			break
		}
	}
	// 查询日志记录
	var orderFlowLast model.OrderFlow
	if err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND del = 0 AND status = 'IssuerAgreeStage'", stage.OrderId).Order("level desc").First(&orderFlowLast).Error; err != nil {
		return err
	}
	// 修改JSON
	newValue, err := sjson.SetRaw(orderFlowLast.Obj, "stages."+strconv.Itoa(index), attachmentNew.Array()[index].Raw)
	if err != nil {
		return err
	}
	orderFlowLast.Obj = newValue
	// 上传JSON
	err, hashJSON := IPFSUploadJSON(newValue)
	if err != nil {
		return err
	}
	orderFlowLast.Attachment = hashJSON
	orderFlowLast.Status = "IssuerAgreeStage" // 修改为IssuerAgreeStage
	// 查询level
	var level int64
	if err = tx.Model(&model.OrderFlow{}).Where("order_id = ?", stage.OrderId).Count(&level).Error; err != nil {
		return err
	}
	orderFlowLast.Level = level + 1 // 节点
	orderFlowLast.Audit = 0
	orderFlowLast.Operator = address
	orderFlowLast.CreatedAt = time.Time{} // 清除
	orderFlowLast.ID = 0                  // 清除
	// 存回orderFlow
	if err = tx.Model(&model.OrderFlow{}).Create(&orderFlowLast).Error; err != nil {
		return err
	}
	return nil
}

// saveFlow 保存日志
func saveFlow(tx *gorm.DB, stage request.UpdatedStageRequest, address string) (err error) {
	var order model.Order
	if err = tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).First(&order).Error; err != nil {
		return err
	}
	// 查询level
	var level int64
	if err = tx.Model(&model.OrderFlow{}).Where("order_id = ?", stage.OrderId).Count(&level).Error; err != nil {
		return err
	}
	// 插入日志表
	orderFlow := model.OrderFlow{OrderId: stage.OrderId, Signature: stage.Signature, SignAddress: stage.SignAddress, SignNonce: stage.SignNonce}
	orderFlow.Level = level + 1 // 节点
	// 阶段状态
	if stage.Status != "" {
		orderFlow.Status = stage.Status
	} else {
		orderFlow.Status = order.Status
	}
	// 阶段详情交付物JSON
	if stage.Obj != "" {
		orderFlow.Obj = stage.Obj
	} else {
		orderFlow.Obj, _ = utils.GetJSONFromCid(order.Attachment)
	}
	// 阶段划分JSON
	if stage.Stages != "" {
		orderFlow.Stages = stage.Stages
	} else {
		orderFlow.Stages = order.Stages
	}
	// JSON IPFS
	if stage.Attachment != "" {
		orderFlow.Attachment = stage.Attachment // JSON IPFS
	} else {
		orderFlow.Attachment = order.Attachment
	}
	orderFlow.Operator = address // 操作人
	if err = tx.Model(&model.OrderFlow{}).Create(&orderFlow).Error; err != nil {
		return err
	}
	return nil
}

// rollbackStatus 回滚操作
func rollbackStatus(tx *gorm.DB, old string, new string, stage request.UpdatedStageRequest) (err error, okRun bool) {
	if old == "WaitIssuerAgree" && new == "InvalidSign" {
		if err = tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).Update("status", "WaitWorkerStage").Error; err != nil {
			return err, true
		}
		return err, true
	}
	if old == "WorkerAgreeStage" && new == "InvalidSign" {
		if err = tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).Update("status", "WaitWorkerConfirmStage").Error; err != nil {
			return err, true
		}
		return err, true
	}
	statusMap := map[string][]string{"WaitProlongAgree": {"DisagreeProlong", "InvalidSign"}, "WaitAppendAgree": {"DisagreeAppend", "InvalidSign"}}
	// 校验阶段状态流转
	status, ok := statusMap[old]
	if !ok {
		return nil, okRun
	}
	if utils.SliceIsExist(status, new) {
		rollMap := map[string]string{"DisagreeProlong": "IssuerAgreeStage", "DisagreeAppend": "IssuerAgreeStage", "InvalidSign": "IssuerAgreeStage"}
		rollStatus, ok := rollMap[new]
		if !ok {
			return nil, okRun
		}
		// 获取回滚信息
		var orderFlow model.OrderFlow
		err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND status = ? AND del = 0", stage.OrderId, rollStatus).Order("level desc").First(&orderFlow).Error
		if err != nil {
			return err, okRun
		}
		// 回滚操作
		order := map[string]interface{}{"attachment": orderFlow.Attachment, "stages": orderFlow.Stages, "Status": orderFlow.Status, "signature": orderFlow.Signature, "sign_address": orderFlow.SignAddress, "sign_nonce": orderFlow.SignNonce}
		raw := tx.Model(&model.Order{}).Where("order_id = ?", stage.OrderId).Updates(&order)
		if raw.RowsAffected == 0 {
			return errors.New("回滚失败"), okRun
		}
		if raw.Error != nil {
			return err, okRun
		}
		// 删除过程信息
		if err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND level > ?", stage.OrderId, orderFlow.Level).Update("del", 1).Error; err != nil {
			return err, okRun
		}
		// 需要进行下一步
		if new == "IssuerAgreeStage" {
			return err, okRun
		}
		return nil, true
	}
	return nil, okRun
}

// statusValid 校验阶段状态流转
func statusValid(old string, new string, stage request.UpdatedStageRequest, order model.Order) (err error) {
	// 状态正确性
	if !utils.SliceIsExist([]string{"WaitWorkerStage", "WaitIssuerAgree", "IssuerAgreeStage", "WaitWorkerConfirmStage",
		"WorkerAgreeStage", "WaitProlongAgree", "AgreeProlong", "DisagreeProlong", "WaitAppendAgree", "AgreeAppend", "DisagreeAppend", "InvalidSign"}, new) {
		return errors.New("状态错误")
	}
	// 需要验证链上状态
	if utils.SliceIsExist([]string{"AgreeProlong", "AgreeAppend"}, new) && stage.Hash == "" {
		return errors.New("hash不能为空")
	}
	// 校验阶段状态流转
	statusMap := map[string][]string{"WaitWorkerStage": {"WaitIssuerAgree"},
		"WaitIssuerAgree":        {"WaitIssuerAgree", "IssuerAgreeStage", "WaitWorkerConfirmStage", "InvalidSign"},
		"WaitWorkerConfirmStage": {"WaitWorkerConfirmStage", "WorkerAgreeStage", "WaitIssuerAgree"},
		"WorkerAgreeStage":       {"WaitIssuerAgree", "IssuerAgreeStage", "WaitWorkerConfirmStage", "InvalidSign"},
		"IssuerAgreeStage":       {"WaitProlongAgree", "WaitAppendAgree", "AbortOrder"},
		"WaitProlongAgree":       {"AgreeProlong", "DisagreeProlong", "InvalidSign"},
		"WaitAppendAgree":        {"AgreeAppend", "DisagreeAppend", "InvalidSign"},
	}
	status, ok := statusMap[old]
	if !ok {
		return errors.New("状态流转不存在")
	}
	if !utils.SliceIsExist(status, new) {
		return errors.New("状态流转错误")
	}
	// 任务已完成
	if order.Progress > 2 {
		return errors.New("任务已结束")
	}
	return nil
}

// storeHash 保存Hash
func storeHash(stage request.UpdatedStageRequest, address string, chain string) (runBool bool, err error) {
	if stage.Hash == "" {
		return runBool, nil
	}
	if stage.Status == "AgreeProlong" {
		runBool = true
		transHash := model.TransHash{SendAddr: address, EventName: "ProlongStage", Hash: stage.Hash}
		if err = SaveHash(transHash, chain); err != nil {
			return runBool, errors.New("操作失败")
		}
	} else if stage.Status == "AgreeAppend" {
		runBool = true
		transHash := model.TransHash{SendAddr: address, EventName: "AppendStage", Hash: stage.Hash}
		if err = SaveHash(transHash, chain); err != nil {
			return runBool, errors.New("操作失败")
		}
	} else if stage.Status == "IssuerAgreeStage" {
		runBool = true
		transHash := model.TransHash{SendAddr: address, EventName: "ConfirmOrderStage", Hash: stage.Hash}
		if err = SaveHash(transHash, chain); err != nil {
			return runBool, errors.New("操作失败")
		}
	} else if stage.Signature != "" {
		runBool = true
		raw, err := json.Marshal(stage)
		if err != nil {
			return runBool, errors.New("操作失败")
		}
		transHash := model.TransHash{SendAddr: address, EventName: "OrderStarted", Hash: stage.Hash, Raw: string(raw)}
		if err = SaveHash(transHash, chain); err != nil {
			return runBool, errors.New("操作失败")
		}
	} else if stage.Obj == "" {
		runBool = true
		transHash := model.TransHash{SendAddr: address, EventName: "", Hash: stage.Hash}
		if err = SaveHash(transHash, chain); err != nil {
			return runBool, errors.New("操作失败")
		}
	}
	return runBool, nil
}

// sendMessage 发送消息
func sendMessage(tx *gorm.DB, order model.Order, stage request.UpdatedStageRequest, orderFlowTop model.OrderFlow, sender string, chain string) (err error) {
	// 查询task
	var task model.Task
	if err = tx.Model(&model.Task{}).Where("task_id =?", order.TaskID).First(&task).Error; err != nil {
		return err
	}
	status := stage.Status
	// 状态改变发送消息
	if status == "WaitIssuerAgree" || status == "WorkerAgreeStage" || status == "WorkerDelivery" || status == "InvalidSign" {
		// 发送消息
		if err = message.Template(status, utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, "", chain); err != nil {
			return err
		}
	}
	// 区分甲方乙方
	if status == "WaitWorkerConfirmStage" {
		// 发送消息
		if err = message.Template(status, utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, sender, chain); err != nil {
			return err
		}
	}
	// Tip
	if order.Status == status && (status == "WaitIssuerAgree" || status == "WaitWorkerConfirmStage") {
		// 发送消息
		if err = message.Template("Tip"+status, utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, sender, chain); err != nil {
			return err
		}
	}
	// tip
	if order.Status == "WaitWorkerConfirmStage" && status == "WaitIssuerAgree" {
		// 发送消息
		if err = message.Template("Tip"+status, utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, sender, chain); err != nil {
			return err
		}
	}

	if status == "DisagreeAppend" {
		mapStr := utils.StructToMap([]any{order, task})
		stagesNew := gjson.Get(orderFlowTop.Obj, "stages.#.milestone.title")
		mapStr["stage_name"] = stagesNew.Array()[len(stagesNew.Array())-1].String()
		// 发送消息
		if err = message.Template(status, mapStr, order.Issuer, order.Worker, sender, chain); err != nil {
			return err
		}
	}
	// 拒绝阶段延长
	if status == "DisagreeProlong" {
		// 查询日志记录
		var orderFlowLast model.OrderFlow
		if err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND del = 0 AND status = 'IssuerAgreeStage'", stage.OrderId).Order("level desc").First(&orderFlowLast).Error; err != nil {
			return err
		}
		stagesNew := gjson.Get(order.Stages, "period")
		stagesOld := gjson.Get(orderFlowLast.Stages, "period")
		// 阶段数量相同
		if len(stagesNew.Array()) == len(stagesOld.Array()) {
			for i := 0; i < len(stagesNew.Array()); i++ {
				if stagesNew.Array()[i].Int() != stagesOld.Array()[i].Int() {
					// 如果有预付款
					if stagesNew.Array()[0].Int() == 0 {
						i -= 1
					}
					mapStr := utils.StructToMap([]any{order, task})
					mapStr["stage"] = "P" + strconv.Itoa(i+1)
					// 发送消息
					if err = message.Template(status, mapStr, order.Issuer, order.Worker, sender, chain); err != nil {
						return err
					}
					break
				}
			}
		}
	}

	// 申请阶段延长
	if status == "WaitProlongAgree" && gjson.Get(orderFlowTop.Stages, "period").String() != gjson.Get(stage.Stages, "period").String() {
		// 判断是否阶段延长
		stagesNew := gjson.Get(stage.Stages, "period")
		stagesOld := gjson.Get(orderFlowTop.Stages, "period")
		// 阶段数量相同
		if len(stagesNew.Array()) == len(stagesOld.Array()) {
			for i := 0; i < len(stagesNew.Array()); i++ {
				if stagesNew.Array()[i].Int() != stagesOld.Array()[i].Int() {
					// 如果有预付款
					if stagesNew.Array()[0].Int() == 0 {
						i -= 1
					}
					mapStr := utils.StructToMap([]any{order, task})
					mapStr["stage"] = "P" + strconv.Itoa(i+1)
					// 发送消息
					if err = message.Template(status, mapStr, order.Issuer, order.Worker, sender, chain); err != nil {
						return err
					}
					break
				}
			}
		}
	}
	// 申请添加阶段
	if status == "WaitAppendAgree" {
		fmt.Println(gjson.Get(order.Stages, "period").String())
		fmt.Println(gjson.Get(stage.Stages, "period").String())
		if gjson.Get(order.Stages, "period").String() != gjson.Get(stage.Stages, "period").String() {
			stagesNew := gjson.Get(stage.Obj, "stages.#.milestone.title")
			mapStr := utils.StructToMap([]any{order, task})
			mapStr["stage_name"] = stagesNew.Array()[len(stagesNew.Array())-1].String()
			// 发送消息
			if err = message.Template(status, mapStr, order.Issuer, order.Worker, sender, chain); err != nil {
				return err
			}
		} else {
			stagesNew := gjson.Get(stage.Obj, "stages.#.milestone.title")
			mapStr := utils.StructToMap([]any{order, task})
			mapStr["stage_name"] = stagesNew.Array()[len(stagesNew.Array())-1].String()
			if err = message.Template("WaitAppendPayment", mapStr, order.Issuer, order.Worker, "", chain); err != nil {
				return err
			}
		}
	}
	// 阶段提交交付
	if stage.Obj != "" && gjson.Get(orderFlowTop.Obj, "stages").String() != gjson.Get(stage.Obj, "stages").String() {
		// 判断是否有新交付物
		attachmentNew := gjson.Get(stage.Obj, "stages.#.delivery")
		attachmentOld := gjson.Get(orderFlowTop.Obj, "stages.#.delivery")
		stagesNew := gjson.Get(order.Stages, "period")
		// 不是阶段交付
		if len(attachmentNew.Array()) != len(attachmentOld.Array()) {
			return nil
		}
		// 遍历得到哪个阶段交付
		for i := 0; i < len(attachmentNew.Array()); i++ {
			if attachmentNew.Array()[i].String() != attachmentOld.Array()[i].String() && attachmentNew.Array()[i].String() != "" {
				// 如果有预付款
				if stagesNew.Array()[0].Int() == 0 {
					i -= 1
				}
				mapStr := utils.StructToMap([]any{order, task})
				mapStr["stage"] = "P" + strconv.Itoa(i+1)
				if err = message.Template("WorkerDelivery", mapStr, order.Issuer, order.Worker, "", chain); err != nil {
					return err
				}
				break
			}
		}
	}
	return err
}
