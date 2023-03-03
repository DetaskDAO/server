package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

var i18nMap map[string]string

func init() {
	i18nMap = map[string]string{
		"操作失败":     "Operation failed",
		"操作成功":     "Successful operation",
		"获取成功":     "Achieve success",
		"获取失败":     "Failed to get",
		"上传成功":     "Upload successful",
		"上传失败":     "Upload failed",
		"文件大小超过限制": "File size exceeds limit",
		"文件格式不正确":  "Wrong file format",
	}

}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, i18n(c, "操作成功"), c)
}

func OkWithRaw(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, i18n(c, "操作成功"), c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, i18n(c, message), c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, i18n(c, "操作失败"), c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func i18n(c *gin.Context, msg string) string {
	if c.GetString("lang") != "zh" {
		if v, ok := i18nMap[msg]; ok {
			return v
		}
	}
	return msg
}
