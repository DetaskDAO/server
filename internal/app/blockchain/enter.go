package blockchain

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// HandleTraverseFailed 交易失败处理
func HandleTraverseFailed(chain string, transHash string, status uint) error {
	// 删除任务
	if err := global.MAPDB[chain].Model(&model.TransHash{}).Where("hash = ?", transHash).Updates(map[string]interface{}{"raw": "", "status": status, "deleted_at": time.Now()}).Error; err != nil {
		return err
	}
	return nil
}

// EventsParser 事件解析处理
func EventsParser(chain string, transHash model.TransHash, Logs []*types.Log) (err error) {
	for _, vLog := range Logs {
		name, ok := global.ContractEvent[vLog.Topics[0]]
		if !ok {
			continue
		}
		fmt.Println(name)
		switch name {
		case "OrderCreated":
			err = ParseOrderCreated(chain, vLog)
			return err
		case "OrderAbort":
			err = ParseOrderAbort(chain, vLog)
			return err
		case "ConfirmOrderStage":
			err = parseConfirmOrderStage(chain, vLog)
			return err
		case "AppendStage":
			err = parseAppendStage(chain, transHash, vLog)
			return err
		case "ProlongStage":
			err = parseProlongStage(chain, transHash, vLog)
			return err
		case "TaskCreated":
			err = ParseTaskCreated(chain, vLog)
			return err
		case "TaskModified":
			err = ParseTaskModified(chain, vLog)
			return err
		case "TaskDisabled":
			err = ParseTaskDisabled(chain, vLog)
			return err
		case "ApplyFor":
			err = ParseApplyFor(chain, transHash, vLog)
			return err
		case "CancelApply":
			err = ParseCancelApply(chain, vLog)
			return err
		case "OrderStarted":
			err = ParseOrderStarted(chain, transHash, vLog)
		case "Withdraw":
			err = ParseWithdraw(chain, vLog)
		default:
			fmt.Printf("Unknown")
		}
	}
	HandleTraverseFailed(chain, transHash.Hash, 4)
	return nil
}
