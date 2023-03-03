package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/model"
	"code-market-admin/internal/app/utils"
	"github.com/tidwall/gjson"
	"strings"
)

var stats model.Stats

func GetStats() model.Stats {
	return stats
}

func UpdateStats() {
	var temp model.Stats
	// 获取当前 人数
	for _, v := range global.MAPDB {
		var count int64
		err := v.Model(&model.User{}).Count(&count).Error
		if err != nil {
			continue
		}
		temp.TotalMembers += count
	}
	payload := strings.NewReader("{\"query\":\"{\\n  taskCounter(id: \\\"1\\\")  {\\n    TaskCreatedCount\\n  }\\n  orderNFTCounter(id: \\\"1\\\") {\\n    OrderNFTCount\\n  }\\n}\"}")

	body, err := utils.GraphQlRequest("https://api.thegraph.com/subgraphs/name/liangjies/detask", payload)
	if err != nil {
		return
	}
	if gjson.Valid(string(body)) {
		temp.TotalTask = gjson.Get(string(body), "data.taskCounter.TaskCreatedCount").Uint()
		temp.TotalNFT = gjson.Get(string(body), "data.orderNFTCounter.OrderNFTCount").Uint()*2 + temp.TotalTask
	}

	if temp.TotalMembers > stats.TotalMembers {
		stats.TotalMembers = temp.TotalMembers
	}
	if temp.TotalTask > stats.TotalTask {
		stats.TotalTask = temp.TotalTask
	}
	if temp.TotalNFT > stats.TotalNFT {
		stats.TotalNFT = temp.TotalNFT
	}
}
