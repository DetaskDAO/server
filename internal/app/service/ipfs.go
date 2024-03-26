package service

import (
	"code-market-admin/internal/app/global"
	"code-market-admin/internal/app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
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
	// 保存 IPFS
	f, openError := header.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("file.Open() Filed", zap.Error(err))
		return err, hash
	}
	// 目录
	director := global.CONFIG.Local.IPFS + "/"
	mkdirErr := os.MkdirAll(director, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("os.MkdirAll() Filed", zap.Error(err))
	}
	p := director + resJson.Hash
	out, createErr := os.Create(p)
	if createErr != nil {
		global.LOG.Error("os.Create() Filed", zap.Error(err))
		return err, hash
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("io.Copy() Filed", zap.Error(err))
		return err, hash
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
	director := global.CONFIG.Local.IPFS + "/"
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		global.LOG.Error("os.MkdirAll() Filed", zap.Error(err))
		return err, hash
	}
	p := director + resJson.Hash
	err = os.WriteFile(p, []byte(data), 0664)
	if err != nil {
		global.LOG.Error("os.WriteFile() Filed", zap.Error(err))
		return err, hash
	}
	return err, resJson.Hash
}
