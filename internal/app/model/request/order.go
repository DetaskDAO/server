package request

import "code-market-admin/internal/app/model"

type GetOrderListRequest struct {
	model.Order
	PageInfo
}
type CreateOrderRequest struct {
	Hash string `json:"hash" form:"hash"` // 交易hash
	model.Order
}
type DeleteOrderRequest struct {
	model.Order
}
type UpdatedStageRequest struct {
	model.Order
	Hash string `json:"hash" form:"hash"` // 交易hash
	Obj  string `json:"obj" form:"obj"`
}

type UpdatedProgressRequest struct {
	model.Order
}
