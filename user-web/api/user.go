package api

import (
	"go.uber.org/zap"
	"micro/user-web/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

}
