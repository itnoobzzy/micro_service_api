package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"micro/demo/token_test/middlewares"
	"net/http"
	"time"
)

func GetToken(c *gin.Context) {
	j := middlewares.NewJwt()
	claims := middlewares.CustomClaims{
		// payload信息, 应该从请求参数中获取
		ID:          uint(123),
		NickName:    "zhouzy",
		AuthorityId: uint(12345),
		// 生成JWT token 所需信息
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "签名机构",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         uint(123),
		"nick_name":  "zhouzy",
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func NeedToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func main() {
	r := gin.Default()
	// 获取token
	r.GET("/token", GetToken)
	// 请求头要有x-token 值为token
	r.GET("/need_token", middlewares.JWTAuth(), NeedToken)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
