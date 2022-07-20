package initialize

import (
	"github.com/gin-gonic/gin"
	"micro/goods-web/router"
	"net/http"

	"micro/goods-web/middlewares"
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
	router.InitGoodsRouter(ApiGroup)
	//router.InitCategoryRouter(ApiGroup)
	//router.InitBannerRouter(ApiGroup)
	//router.InitBrandRouter(ApiGroup)

	return Router
}
