package utils

import (
	"code-market-admin/internal/app/global"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"sync/atomic"
	"time"
)

var index int
var lock atomic.Bool

func GetIPFSUploadAPI() string {
	return global.CONFIG.IPFS[index].UploadAPI
}

func GetIPFSGateway(hash string) string {
	if hash == "" {
		return global.CONFIG.IPFS[index].API
	}
	url := fmt.Sprintf("%s/%s", global.CONFIG.IPFS[index].API, hash)
	return url
}

func BalanceIPFS() {
	if lock.Load() {
		return
	}
	lock.Store(true)
	defer lock.Store(false)

	IPFS := global.CONFIG.IPFS
	indexList := make([]int64, len(IPFS))
	for i, v := range IPFS {
		if v.API == "" || v.UploadAPI == "" {
			return
		}
		spent, err := ipfsRequest(v.API, v.UploadAPI)
		if err != nil {
			fmt.Println(err)
		}
		indexList[i] = spent
		time.Sleep(time.Second * 1)
	}
	fmt.Println(indexList)
	index, _ = SliceMin[int64](indexList)
	global.LOG.Warn("IPFS 切换: " + strconv.Itoa(index))
}

func ipfsRequest(api string, uploadAPI string) (spent int64, err error) {
	defer func() {
		if err := recover(); err != nil {
			spent = 9999999999999
			return
		}
	}()
	startTime := time.Now()
	// 上传JSON
	// 组成请求体
	jsonReq := make(map[string]interface{})
	jsonReq["body"] = "{\"foo\":\"bar\"}"
	// 发送请求
	url := fmt.Sprintf("%s/upload/json", uploadAPI)
	res, err := PostRequest(url, jsonReq, "application/json")
	if err != nil {
		return 9999999999999, err
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
		return 9999999999999, err
	}
	if resJson.Status != "1" {
		return 9999999999999, err
	}
	// 请求JSON
	urlReq := fmt.Sprintf("%s/%s", api, resJson.Hash)
	content, err := GetRequest(urlReq)
	if err != nil || !gjson.Valid(content) {
		return 9999999999999, err
	}
	return time.Since(startTime).Milliseconds(), nil
}
