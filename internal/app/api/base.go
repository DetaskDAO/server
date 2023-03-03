package api

import (
	"code-market-admin/internal/app/model/response"
	"code-market-admin/internal/app/service"
	"github.com/gin-gonic/gin"
)

// GetStats @Base
// @Summary 获取生态信息
// @Router /base/stats [get]
func GetStats(c *gin.Context) {
	response.OkWithRaw(service.GetStats(), c)
}
