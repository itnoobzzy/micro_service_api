package router

import (
	"github.com/gin-gonic/gin"
	"micro/order-web/api/order"
	"micro/order-web/api/pay"
	"micro/order-web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	//OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
