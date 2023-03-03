package blockchain

import (
	ABI "code-market-admin/abi"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/message"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var deOrderAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.DeOrderMetaData.ABI))
	if err != nil {
		global.LOG.Error("Failed to Load Abi", zap.Error(err))
		panic(err)
	}
	deOrderAbi = contractAbi
}

// ParseOrderCreated 解析OrderCreated事件
func ParseOrderCreated(chain string, vLog *types.Log) (err error) {
	var orderCreated ABI.DeOrderOrderCreated
	ParseErr := deOrderAbi.UnpackIntoInterface(&orderCreated, "OrderCreated", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	order := model.Order{TaskID: vLog.Topics[1].Big().Int64(), OrderId: vLog.Topics[2].Big().Int64()}
	order.Issuer = orderCreated.Issuer.String() // 甲方
	order.Worker = orderCreated.Worker.String() // 乙方
	// 解析 币种
	order.Currency = orderCreated.Token.String()
	order.Amount = orderCreated.Amount.String()
	var orderStruct model.Task
	if errOrder := tx.Model(&model.Task{}).Where("task_id = ?", order.TaskID).First(&orderStruct).Error; errOrder != nil {
		if errOrder != gorm.ErrRecordNotFound {
			tx.Rollback()
			return errOrder
		}
	}
	orderByte, _ := json.Marshal(orderStruct)
	order.Task = string(orderByte)
	// 更新||插入数据
	err = tx.Model(&model.Order{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}},
		UpdateAll: true,
	}).Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 插入日志表
	orderFlow := model.OrderFlow{OrderId: order.OrderId}
	err = tx.Model(&model.OrderFlow{}).Create(&orderFlow).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新apply表状态
	_ = tx.Model(&model.Apply{}).Where("apply_addr = ? AND task_id = ?", order.Worker, order.TaskID).Updates(map[string]interface{}{"status": 2, "order_id": order.OrderId}).Error
	// 发送消息
	var task model.Task
	_ = tx.Model(&model.Task{}).Where("task_id = ?", order.TaskID).First(&task).Error
	if err = message.Template("OrderCreated", utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, "", chain); err != nil {
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// ParseOrderStarted 解析OrderStarted事件
func ParseOrderStarted(chain string, transHash model.TransHash, vLog *types.Log) (err error) {
	var orderStarted ABI.DeOrderOrderStarted
	ParseErr := deOrderAbi.UnpackIntoInterface(&orderStarted, "OrderStarted", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	orderId := vLog.Topics[1].Big().Int64()
	// 解析Raw数据
	signature := gjson.Get(transHash.Raw, "signature").String()
	// 正常情况
	raw := tx.Model(&model.Order{}).Unscoped().Where("order_id = ? AND signature = ?", orderId, signature).Updates(map[string]interface{}{"signature": "", "sign_address": "", "sign_nonce": 0, "status": "IssuerAgreeStage", "progress": 2, "deleted_at": nil})
	// 异常情况
	if raw.RowsAffected == 0 {
		// 获取回滚信息
		var orderFlow model.OrderFlow
		err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND signature = ? AND (status = 'WaitIssuerAgree' or status = 'WorkerAgreeStage') AND del = 0", orderId, signature).Order("level desc").First(&orderFlow).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		// 回滚操作 && 确认
		order := map[string]interface{}{"attachment": orderFlow.Attachment, "stages": orderFlow.Stages, "Status": "IssuerAgreeStage", "signature": "", "sign_address": "", "sign_nonce": 0, "progress": 2, "deleted_at": nil}
		raw := tx.Model(&model.Order{}).Unscoped().Where("order_id = ?", orderId).Updates(&order)
		if raw.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("操作失败")
		}
	}
	if raw.Error != nil {
		tx.Rollback()
		return raw.Error
	}
	// 插入日志表
	if errSave := saveOrderFlow(tx, orderId); errSave != nil {
		global.LOG.Error("SaveOrderFlow error: ", zap.Error(errSave))
	}
	// 发送消息
	_ = sendMessage(chain, uint64(orderId), "OrderStarted", "", nil)
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// UpdatedProgress 更新任务Progress状态
func UpdatedProgress(orderID int64, chain string) (err error) {
	// 错误处理
	defer func() {
		if errRe := recover(); errRe != nil {
			global.LOG.Error("UpdatedProgress错误", zap.Any("err:", err))
			err = errors.New("error")
		}
	}()
	// 获取当前progress
	var progress uint8
	if err = global.MAPDB[chain].Model(&model.Order{}).Select("progress").Where("order_id = ?", orderID).First(&progress).Error; err != nil {
		return err
	}
	// client
	client, err := ethclient.Dial(global.ProviderMap[chain])
	if err != nil {
		return err
	}
	// 合约地址
	address := global.ContractAddr[chain+":DeOrder"]
	instance, err := ABI.NewDeOrder(address, client)
	if err != nil {
		return err
	}
	order, err := instance.GetOrder(nil, big.NewInt(orderID))
	if err != nil {
		return err
	}
	// 获取失败
	if order.TaskId.Int64() == 0 && progress != 0 {
		return errors.New("操作失败")
	}
	// 修改Amount
	if err = global.MAPDB[chain].Model(&model.Order{}).Where("order_id = ?", orderID).Updates(map[string]interface{}{"amount": order.Amount.String(), "pay_type": order.PayType}).Error; err != nil {
		return err
	}
	// 任务完成 修改state
	if (order.Progress == 3 && progress != 3) || (order.Progress == 4 && progress != 4) || (order.Progress == 5 && progress != 5) {
		if err = orderDoneOperation(chain, orderID); err != nil {
			return err
		}
	}
	// 状态更新
	if order.Progress != progress {
		// 更新progress
		raw := global.MAPDB[chain].Model(&model.Order{}).Where("order_id = ?", orderID).Update("progress", order.Progress)
		if raw.RowsAffected == 0 {
			return errors.New("操作失败")
		}
		// 发送消息
		if err = SendMessage(chain, model.Order{TaskID: order.TaskId.Int64(), Issuer: order.Issuer.String(), Worker: order.Worker.String(), OrderId: orderID, Progress: order.Progress}); err != nil {
			return err
		}
		return raw.Error
	}
	return nil
}

// ParseOrderAbort 解析OrderAbort事件
func ParseOrderAbort(chain string, vLog *types.Log) (err error) {
	var orderAbort ABI.DeOrderOrderAbort
	ParseErr := deOrderAbi.UnpackIntoInterface(&orderAbort, "OrderAbort", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新任务状态
	orderID := vLog.Topics[1].Big().Int64()
	operation := func() error {
		return UpdatedProgress(orderID, chain)
	}
	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)); err != nil {
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// ParseWithdraw 解析Withdraw事件
func ParseWithdraw(chain string, vLog *types.Log) (err error) {
	var withdraw ABI.DeOrderWithdraw
	ParseErr := deOrderAbi.UnpackIntoInterface(&withdraw, "Withdraw", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新任务状态
	orderID := vLog.Topics[1].Big().Int64()
	operation := func() error {
		for i := 0; i < 3; i++ {
			err1 := UpdatedProgress(orderID, chain)
			time.Sleep(1 * time.Second)
			if err1 != nil {
				return err
			}
		}
		return nil
	}
	if err = operation(); err != nil {
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// SendMessage 发送消息
func SendMessage(chain string, order model.Order) (err error) {
	var task model.Task
	if err = global.MAPDB[chain].Model(&model.Task{}).Where("task_id = ?", order.TaskID).First(&task).Error; err != nil {
		return err
	}
	if order.Progress == 5 {
		// 任务完成
		if err = message.Template("OrderDone", utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, "", chain); err != nil {
			return err
		}
	} else if order.Progress == 2 {
		// 任务开始

	} else if order.Progress == 3 {
		// 甲方中止任务
		if err = message.Template("OrderAbort", utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, order.Issuer, chain); err != nil {
			return err
		}
	} else if order.Progress == 4 {
		// 乙方中止任务
		if err = message.Template("OrderAbort", utils.StructToMap([]any{order, task}), order.Issuer, order.Worker, order.Worker, chain); err != nil {
			return err
		}
	}
	return nil
}

// orderDoneOperation 状态操作
func orderDoneOperation(chain string, orderID int64) (err error) {
	// 修改任务状态
	tx := global.MAPDB[chain].Begin()
	// 查询任务
	var order model.Order
	if err = tx.Model(&model.Order{}).Where("order_id =?", orderID).First(&order).Error; err != nil {
		tx.Rollback()
		return err
	}
	updateOrder := map[string]interface{}{"state": 1, "pending": 0}
	if order.Status != "IssuerAgreeStage" {
		// 获取回滚信息
		var orderFlow model.OrderFlow
		err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND status = ? AND del = 0", order.OrderId, "IssuerAgreeStage").Order("level desc").First(&orderFlow).Error
		if err != nil {
			return err
		}
		//
		updateOrder["signature"] = ""
		updateOrder["sign_address"] = ""
		updateOrder["stages"] = orderFlow.Stages
		updateOrder["attachment"] = orderFlow.Attachment
		updateOrder["status"] = "IssuerAgreeStage"
	}
	raw := tx.Model(&model.Order{}).Where("order_id = ?", orderID).Updates(updateOrder)
	if raw.RowsAffected == 0 || raw.Error != nil {
		tx.Rollback()
		return errors.New("操作失败")
	}

	rawApply := tx.Model(&model.Apply{}).Where("order_id = ?", orderID).Update("Status", 1)
	if rawApply.Error != nil {
		tx.Rollback()
		return errors.New("操作失败")
	}
	if err := saveOrderFlow(tx, orderID); err != nil {
		global.LOG.Error("Error saving order flow")
		tx.Rollback()
		return errors.New("操作失败")
	}
	// 删除apply信息
	//if err = tx.Model(&model.Apply{}).Unscoped().Where("task_id =? AND apply_addr = ?", order.TaskID, order.Worker).Delete(&model.Apply{}).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}

	return tx.Commit().Error
}

// issuerAgreeOperation 状态操作
func issuerAgreeOperation(chain string, orderID int64) (err error) {
	// 清空签名 && 修改状态
	tx := global.MAPDB[chain].Begin()
	raw := tx.Model(&model.Order{}).Where("order_id = ?", orderID).Updates(map[string]interface{}{"signature": "", "sign_address": "", "sign_nonce": 0, "status": "IssuerAgreeStage"})
	if raw.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("操作失败")
	}
	if err = saveOrderFlow(tx, orderID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func saveOrderFlow(tx *gorm.DB, orderID int64) (err error) {
	// 查询当前记录
	var order model.Order
	if err = tx.Model(&model.Order{}).Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return err
	}
	// 查询level
	var level int64
	if err = tx.Model(&model.OrderFlow{}).Where("order_id = ?", order.OrderId).Count(&level).Error; err != nil {
		return err
	}
	// 插入日志表
	orderFlow := model.OrderFlow{OrderId: order.OrderId, Status: order.Status, Stages: order.Stages}
	orderFlow.Level = level + 1             // 节点
	orderFlow.Attachment = order.Attachment // JSON IPFS
	orderFlow.Operator = order.Issuer       // 甲方
	if order.Attachment != "" {
		attachment := order.Attachment
		// 从缓存获取
		orderFlow.Obj, _ = utils.GetJSONFromCid(attachment)
	}
	if err = tx.Model(&model.OrderFlow{}).Create(&orderFlow).Error; err != nil {
		return err
	}

	return nil
}

// ParseConfirmOrderStage 解析ConfirmOrderStage事件
func parseConfirmOrderStage(chain string, vLog *types.Log) (err error) {
	var confirmOrderStage ABI.DeOrderConfirmOrderStage
	ParseErr := deOrderAbi.UnpackIntoInterface(&confirmOrderStage, "ConfirmOrderStage", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	orderID := vLog.Topics[1].Big().Uint64()
	// 发送消息
	type StageData struct {
		Stage string `json:"stage"`
	}
	if err = sendMessage(chain, orderID, "ConfirmOrderStage", "", StageData{Stage: "P" + confirmOrderStage.StageIndex.String()}); err != nil {
		tx.Rollback()
		return err
	}
	// update withdraw balance
	operation := func() error {
		return UpdatedPendingWithdraw(int64(orderID), chain)
	}
	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3)); err != nil {
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// parseAppendStage 解析AppendStage事件
func parseAppendStage(chain string, transHash model.TransHash, vLog *types.Log) (err error) {
	var appendStage ABI.DeOrderAppendStage
	ParseErr := deOrderAbi.UnpackIntoInterface(&appendStage, "AppendStage", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	orderID := vLog.Topics[1].Big().Uint64()
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND audit = 0 AND del = 0", orderID).Update("audit", 1).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 保存order日志 && 将Order状态改变
	if err = issuerAgreeOperation(chain, int64(orderID)); err != nil {
		tx.Rollback()
		return err
	}
	// 发送消息
	if err = sendMessage(chain, orderID, "AgreeAppend", transHash.SendAddr, nil); err != nil {
		tx.Rollback()
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// parseProlongStage 解析ProlongStage事件
func parseProlongStage(chain string, transHash model.TransHash, vLog *types.Log) (err error) {
	var prolongStage ABI.DeOrderProlongStage
	ParseErr := deOrderAbi.UnpackIntoInterface(&prolongStage, "ProlongStage", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	orderID := vLog.Topics[1].Big().Uint64()
	// 保存order日志
	if err = issuerAgreeOperation(chain, int64(orderID)); err != nil {
		return err
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	err = tx.Model(&model.OrderFlow{}).Where("order_id = ? AND audit = 0 AND del = 0", orderID).Update("audit", 1).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 任务状态改变
	err = tx.Model(&model.Order{}).Where("order_id = ?", orderID).Update("status", "IssuerAgreeStage").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 发送消息
	type StageData struct {
		Stage string `json:"stage"`
	}
	if err = sendMessage(chain, orderID, "AgreeProlong", transHash.SendAddr, StageData{Stage: "P" + strconv.Itoa(int(prolongStage.StageIndex.Int64()+1))}); err != nil {
		tx.Rollback()
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// sendMessage 发送消息
func sendMessage(chain string, orderID uint64, status string, sender string, inter any) (err error) {
	// Order信息
	var order model.Order
	if err = global.MAPDB[chain].Model(&model.Order{}).Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return err
	}
	if status == "AgreeAppend" {
		// 查询日志记录
		var orderFlowTop model.OrderFlow
		if err = global.MAPDB[chain].Model(&model.OrderFlow{}).Where("order_id = ? AND del = 0", orderID).Order("level desc").First(&orderFlowTop).Error; err != nil {
			return err
		}
		stagesNew := gjson.Get(orderFlowTop.Obj, "stages.#.milestone.title")
		type StageData struct {
			StageName string `json:"stage_name"`
		}
		stageName := stagesNew.Array()[len(stagesNew.Array())-1].String()
		inter = StageData{StageName: stageName}
	}
	// Task信息
	var task model.Task
	if err = global.MAPDB[chain].Model(&model.Task{}).Where("task_id = ?", order.TaskID).First(&task).Error; err != nil {
		return err
	}
	if err = message.Template(status, utils.StructToMap([]any{order, task, inter}), order.Issuer, order.Worker, sender, chain); err != nil {
		return err
	}
	return nil
}

// UpdatedPendingWithdraw 更新任务Progress状态
func UpdatedPendingWithdraw(orderID int64, chain string) (err error) {
	fmt.Println("UpdatedPendingWithdraw")
	// 错误处理
	defer func() {
		if errRe := recover(); errRe != nil {
			global.LOG.Error("UpdatedPendingWithdraw error", zap.Any("err:", err))
			err = errors.New("error")
		}
	}()
	// client
	client, err := ethclient.Dial(global.ProviderMap[chain])
	if err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		// 合约地址
		address := global.ContractAddr[chain+":DeOrder"]
		instance, err := ABI.NewDeOrder(address, client)
		if err != nil {
			return err
		}
		order, err := instance.PendingWithdraw(nil, big.NewInt(orderID))
		if err != nil {
			return err
		}
		fmt.Println(order)
		if order.NextStage.Uint64() == 0 {
			return nil
		}

		// 修改Amount
		if err = global.MAPDB[chain].Model(&model.Order{}).Where("order_id = ?", orderID).Update("pending", order.Pending.Uint64()).Error; err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func UpdateAllPendingWithdraw() (err error) {
	// 错误处理
	defer func() {
		if errRe := recover(); errRe != nil {
			global.LOG.Error("UpdateAllPendingWithdraw error", zap.Any("err:", errRe))
		}
	}()
	for chain, _ := range global.MAPDB {
		db := global.MAPDB[chain]
		var order []model.Order
		if err = db.Model(&model.Order{}).Where("progress = 2").Find(&order).Error; err != nil {
			return err
		}
		client, err := ethclient.Dial(global.ProviderMap[chain])
		address := global.ContractAddr[chain+":DeOrder"]
		instance, err := ABI.NewDeOrder(address, client)
		if err != nil {
			fmt.Println(err)
			return err
		}
		for _, v := range order {
			pendingWithdraw, errPw := instance.PendingWithdraw(nil, big.NewInt(v.OrderId))
			if errPw != nil {
				if errPw.Error() == "execution reverted" {
					UpdatedProgress(v.OrderId, chain) // Update progress
				}
				continue
			}
			if pendingWithdraw.NextStage.Uint64() == 0 || pendingWithdraw.Pending.Uint64() == uint64(v.Pending) {
				continue
			}
			// 修改Amount
			if err = db.Model(&model.Order{}).Where("order_id = ?", v.OrderId).Update("pending", pendingWithdraw.Pending.Uint64()).Error; err != nil {
				return err
			}
			if pendingWithdraw.Pending.Uint64() > uint64(v.Pending) {
				var task model.Task
				_ = json.Unmarshal([]byte(v.Task), &task)
				// 发送消息
				if errSd := message.Template("PendingWithdraw", utils.StructToMap([]any{v, task}), v.Issuer, v.Worker, "", chain); errSd != nil {
					continue
				}
			}
		}
	}
	return nil
}
