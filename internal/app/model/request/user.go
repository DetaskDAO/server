package request

import "code-market-admin/internal/app/model"

type GetUserInfoRequest struct {
	model.User
}
type UpdateUserInfoRequest struct {
	model.User
}
type CreateUserInfoRequest struct {
	model.User
}

type MsgListRequest struct {
	model.Message
	PageInfo
}

type UnReadRequest struct {
	model.Message
}

type ReadMsgRequest struct {
	model.Message
}
