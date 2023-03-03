package blockchain

import (
	ABI "code-market-admin/abi"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/message"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

var deTaskAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.DeTaskMetaData.ABI))
	if err != nil {
		global.LOG.Error("Failed to Load Abi", zap.Error(err))
		panic(err)
	}
	deTaskAbi = contractAbi
}

// ParseTaskCreated 解析TaskCreated事件
func ParseTaskCreated(chain string, vLog *types.Log) (err error) {
	var taskCreated ABI.DeTaskTaskCreated
	ParseErr := deTaskAbi.UnpackIntoInterface(&taskCreated, "TaskCreated", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	task := model.Task{TaskID: vLog.Topics[1].Big().Uint64(), Title: taskCreated.Task.Title, Period: taskCreated.Task.Period, Attachment: taskCreated.Task.Attachment}
	task.Issuer = taskCreated.Issuer.String()
	task.Budget = taskCreated.Task.Budget.String()
	task.Attachment = taskCreated.Task.Attachment
	task.Currency = utils.CurrencyNames[taskCreated.Task.Currency]
	task.Role = utils.ParseSkills(taskCreated.Task.Skills.Int64())
	// 更新||插入数据
	err = tx.Model(&model.Task{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "task_id"}},
		UpdateAll: true,
	}).Create(&task).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除任务
	if err = tx.Model(&model.TransHash{}).Where("hash = ?", vLog.TxHash.String()).Updates(map[string]interface{}{"raw": "", "status": 1, "deleted_at": time.Now()}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 发送消息
	if err = message.Template("TaskCreated", utils.StructToMap([]any{task}), task.Issuer, "", "", chain); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// ParseTaskModified 解析TaskModified事件
func ParseTaskModified(chain string, vLog *types.Log) (err error) {
	var taskModified ABI.DeTaskTaskModified
	ParseErr := deTaskAbi.UnpackIntoInterface(&taskModified, "TaskModified", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据

	task := model.Task{TaskID: vLog.Topics[1].Big().Uint64(), Title: taskModified.Task.Title, Period: taskModified.Task.Period, Attachment: taskModified.Task.Attachment}
	task.Issuer = taskModified.Issuer.String()
	task.Budget = taskModified.Task.Budget.String()
	task.Attachment = taskModified.Task.Attachment
	task.Currency = utils.CurrencyNames[taskModified.Task.Currency]
	task.Role = utils.ParseSkills(taskModified.Task.Skills.Int64())
	// 更新
	err = tx.Where("task_id", task.TaskID).Updates(&task).Error
	if err != nil {
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

// ParseTaskDisabled 解析TaskDisabled事件
func ParseTaskDisabled(chain string, vLog *types.Log) (err error) {
	var taskDisabled ABI.DeTaskTaskDisabled
	ParseErr := deTaskAbi.UnpackIntoInterface(&taskDisabled, "TaskDisabled", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	if vLog.Topics[2].Big().Int64() == 0 {
		return
	}
	// 删除数据
	taskID := vLog.Topics[1].Big().Uint64()
	raw := tx.Model(&model.Task{}).Where("task_id = ?", taskID).Update("deleted_at", time.Now())
	if raw.Error != nil {
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

// ParseApplyFor 解析ApplyFor事件
func ParseApplyFor(chain string, transHash model.TransHash, vLog *types.Log) (err error) {
	var applyFor ABI.DeTaskApplyFor
	ParseErr := deTaskAbi.UnpackIntoInterface(&applyFor, "ApplyFor", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 更新数据
	taskID := vLog.Topics[1].Big().Uint64()                         // 任务ID
	price := applyFor.Cost.String()                                 // 报价
	applyAddr := common.HexToAddress(vLog.Topics[2].Hex()).String() // 申请人
	apply := model.Apply{TaskID: taskID, Price: price, ApplyAddr: applyAddr}
	// 解析Raw数据
	apply.Desc = gjson.Get(transHash.Raw, "desc").String()
	// 更新||插入数据
	err = tx.Model(&model.Apply{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "apply_addr"}, {Name: "task_id"}},
		UpdateAll: true,
	}).Create(&apply).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 查询Task信息
	var task model.Task
	err = tx.Model(&model.Task{}).Where("task_id =?", taskID).First(&task).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 更新报名人数
	var count int64
	err = tx.Model(&model.Apply{}).Where("task_id =?", taskID).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&model.Task{}).Where("task_id =?", taskID).Update("apply_count", count).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 发送消息
	if err = message.Template("ApplyFor", utils.StructToMap([]any{apply, task}), task.Issuer, apply.ApplyAddr, "", chain); err != nil {
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

// ParseCancelApply 解析CancelApply事件
func ParseCancelApply(chain string, vLog *types.Log) (err error) {
	var cancelApply ABI.DeTaskCancelApply
	ParseErr := deTaskAbi.UnpackIntoInterface(&cancelApply, "CancelApply", vLog.Data)
	if ParseErr != nil {
		return ParseErr
	}
	taskID := vLog.Topics[1].Big().Uint64()  // 任务ID
	applyAddr := cancelApply.Worker.String() // 申请人
	// 开始事务
	tx := global.MAPDB[chain].Begin()
	// 删除数据
	raw := tx.Model(&model.Apply{}).Where("task_id = ?", taskID).Where("apply_addr = ?", applyAddr).Update("deleted_at", time.Now())
	if raw.Error != nil {
		tx.Rollback()
		return
	}
	// 更新报名人数
	var count int64
	err = tx.Model(&model.Apply{}).Where("task_id =?", taskID).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&model.Task{}).Where("task_id =?", taskID).Update("apply_count", count).Error
	if err != nil {
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
