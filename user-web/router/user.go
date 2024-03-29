package router

import (
	"github.com/gin-gonic/gin"
	"micro/user-web/api"
	"micro/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup){
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
	//服务注册和发现
}