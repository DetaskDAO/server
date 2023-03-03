package service

import (
	"code-market-admin/internal/app/blockchain"
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"errors"
)

func SaveHash(transHash model.TransHash, chain string) (err error) {
	transHashRes := global.MAPDB[chain].Model(&model.TransHash{}).Create(&transHash)
	if transHashRes.RowsAffected == 0 {
		return errors.New("新建失败")
	}
	// 启动扫描任务
	_, ok := blockchain.Traversed.Load(chain)
	if !ok {
		go blockchain.HandleTransaction(chain)
	}
	return nil
}
