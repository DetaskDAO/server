package router

import (
	"code-market-admin/internal/app/api"
	"code-market-admin/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order")
	orderRouterWithAuth := Router.Group("order").Use(middleware.JWTAuth())
	{
		orderRouter.GET("getOrderList", api.GetOrderList) // 获取任务列表
	}
	{
		orderRouterWithAuth.POST("updatedStage", api.UpdatedStage)       // 创建阶段划分
		orderRouterWithAuth.POST("updatedProgress", api.UpdatedProgress) // 更新阶段状态
		orderRouterWithAuth.POST("createOrder", api.CreateOrder)         // 创建Order
		orderRouterWithAuth.POST("deleteOrder", api.DeleteOrder)         // 删除Order
	}
}
