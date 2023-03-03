package response

import (
	"code-market-admin/internal/app/model"
)

type GetOrderListResponse struct {
	model.Order
	StageJson     string     `json:"stage_json" form:"stage_json" gorm:"-"`
	LastStageJson string     `json:"last_stage_json" form:"last_stage_json" gorm:"-"`
	LastStages    string     `json:"last_stages" form:"last_stages" gorm:"-"`
	Task          model.Task `json:"task" form:"task"`
	IssuerInfo    user       `json:"issuer_info" form:"issuer_info" gorm:"foreignKey:address;references:issuer"`
	WorkerInfo    user       `json:"worker_info" form:"worker_info" gorm:"foreignKey:address;references:worker"`
}
