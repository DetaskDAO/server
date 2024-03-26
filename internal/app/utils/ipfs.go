package utils

import (
	"code-market-admin/internal/app/global"
	"errors"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

var clientReqIPFS *req.Client
var onceReqIPFS sync.Once

func GetReqIPFSClient() *req.Client {
	onceReqIPFS.Do(func() {
		clientReqIPFS = req.C().SetTimeout(60 * time.Second).SetCommonRetryCount(2)
	})
	return clientReqIPFS
}

func GetJSONFromCid(cid string) (result string, err error) {
	cache, cacheErr := global.JsonCache.Get(cid)
	if cacheErr == nil {
		return string(cache), nil
	}
	// 读取文件内容
	path := global.CONFIG.Local.IPFS
	filePath := path + "/" + cid
	// 判断文件是否存在
	if _, errExist := os.Stat(filePath); os.IsNotExist(errExist) {
		client := GetReqIPFSClient()
		res, err := client.R().Get(GetIPFSGateway(cid))
		if err != nil {
			return result, err
		}
		if !gjson.Valid(res.String()) {
			return result, errors.New("invalid json")
		}
		director := global.CONFIG.Local.IPFS + "/"
		err = os.MkdirAll(director, os.ModePerm)
		if err != nil {
			global.LOG.Error("os.MkdirAll() Filed", zap.Error(err))
			return result, err
		}
		err = os.WriteFile(filePath, res.Bytes(), 0664)
		if err != nil {
			global.LOG.Error("os.WriteFile() Filed", zap.Error(err))
			return result, err
		}
		if err = global.JsonCache.Set(cid, res.Bytes()); err != nil {
			return result, err
		}
		return res.String(), nil
	}
	data, err := os.ReadFile(path + "/" + cid)
	if err != nil {
		global.LOG.Error("读取文件出错: ", zap.Error(err))
		return
	}

	if !gjson.Valid(string(data)) {
		return result, errors.New("invalid json")
	}
	// save cache
	if err = global.JsonCache.Set(cid, data); err != nil {
		return
	}
	return string(data), nil
}
