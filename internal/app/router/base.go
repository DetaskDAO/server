package router

import (
	"code-market-admin/internal/app/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.GET("stats", api.GetStats)
	}
}
