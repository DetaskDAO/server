package service

import (
	"code-market-admin/internal/app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
)

// IPFSUploadFile
// @description: 上传文件
// @param: header *multipart.FileHeader
// @return: err error, list interface{}, total int64
func IPFSUploadFile(header *multipart.FileHeader) (err error, hash string) {
	// 文件大小限制
	if header.Size > 1024*1024*20 {
		return errors.New("文件大小超过限制！"), hash
	}
	// 发送请求
	url := fmt.Sprintf("%s/upload/image", utils.GetIPFSUploadAPI())
	res, err := utils.PostFileRequest(url, header)
	if err != nil {
		go utils.BalanceIPFS()
		return err, hash
	}
	// 解析返回结果
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Hash    string `gorm:"column:hash" json:"hash" form:"hash"`
	}
	var resJson Response
	err = json.Unmarshal(res, &resJson)
	if err != nil {
		return err, hash
	}
	if resJson.Status != "1" {
		go utils.BalanceIPFS()
		return errors.New(resJson.Message), hash
	}

	return err, resJson.Hash
}

// IPFSUploadJSON
// @description: 上传JSON
// @param: header *multipart.FileHeader
// @return: err error, list interface{}, total int64
func IPFSUploadJSON(data string) (err error, hash string) {
	// 组成请求体
	jsonReq := make(map[string]interface{})
	jsonReq["body"] = data
	// 发送请求
	url := fmt.Sprintf("%s/upload/json", utils.GetIPFSUploadAPI())
	res, err := utils.PostRequest(url, jsonReq, "application/json")
	if err != nil {
		go utils.BalanceIPFS()
		return err, hash
	}

	// 解析返回结果
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Hash    string `gorm:"column:hash" json:"hash" form:"hash"`
	}
	var resJson Response
	err = json.Unmarshal(res, &resJson)
	if err != nil {
		return err, hash
	}
	if resJson.Status != "1" {
		go utils.BalanceIPFS()
		return errors.New(resJson.Message), hash
	}
	return err, resJson.Hash
}
