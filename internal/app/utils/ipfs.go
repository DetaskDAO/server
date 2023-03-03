package utils

import (
	"code-market-admin/internal/app/global"
	"errors"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
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
	client := GetReqIPFSClient()
	res, err := client.R().Get(GetIPFSGateway(cid))
	if err != nil {
		return
	}
	if !gjson.Valid(res.String()) {
		return result, errors.New("invalid json")
	}
	// save cache
	if err = global.JsonCache.Set(cid, res.Bytes()); err != nil {
		return
	}
	return res.String(), nil
}
