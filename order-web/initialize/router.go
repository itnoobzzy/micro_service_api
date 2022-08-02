package initialize

import (
	"github.com/gin-gonic/gin"
	"micro/order-web/router"
	"net/http"

	"micro/order-web/middlewares"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())
	//添加链路追踪
	ApiGroup := Router.Group("/g/v1")
	ApiGroup.Use(gin.Logger())
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)

	return Router
}
