package api

import (
	"context"
	"fmt"
	"math/rand"
	"micro/user-web/client"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"micro/user-web/forms"
	"micro/user-web/global"
)

func GenerateSmsCode(witdh int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	//client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", order_global.ServerConfig.AliSmsInfo.ApiKey, order_global.ServerConfig.AliSmsInfo.ApiSecrect)
	//if err != nil {
	//	panic(err)
	//}
	smsCode := GenerateSmsCode(6)
	//request := requests.NewCommonRequest()
	//request.Method = "POST"
	//request.Scheme = "https" // https | http
	//request.Domain = "dysmsapi.aliyuncs.com"
	//request.Version = "2017-05-25"
	//request.ApiName = "SendSms"
	//request.QueryParams["RegionId"] = "cn-beijing"
	//request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile            //手机号
	//request.QueryParams["SignName"] = "慕学在线"                            //阿里云验证过的项目名 自己设置
	//request.QueryParams["TemplateCode"] = "SMS_181850725"               //阿里云的短信模板号 自己设置
	//request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	//response, err := client.ProcessCommonRequest(request)
	//fmt.Print(client.DoAction(request, response))
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
	//将验证码保存起来 - redis
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     fmt.Sprintf("%s:%d", order_global.ServerConfig.RedisInfo.Host, order_global.ServerConfig.RedisInfo.Port),
	//	Password: order_global.ServerConfig.RedisInfo.Password,
	//	DB:       order_global.ServerConfig.RedisInfo.DB,
	//})
	client.Rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
