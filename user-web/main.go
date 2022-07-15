package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"micro/user-web/global"
	"micro/user-web/initialize"
	"micro/user-web/utils"
	"micro/user-web/utils/register/consul"
	myvalidator "micro/user-web/validator"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()

	//2. 初始化配置文件
	initialize.InitConfig()

	//3. 初始化routers
	Router := initialize.Routers()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5. 初始化srv的连接
	initialize.InitSrvConn()
	//6. 初始化redis 连接
	err := initialize.InitRdbClient()
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	debug := viper.GetBool("DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Named("demo").Panic("服务注册失败:", err.Error())
	}

	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Named("demo").Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Named("demo").Panic("启动失败:", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal)
	// 前台时，按 ^C 时触发
	signal.Notify(quit, syscall.SIGINT)
	// 后台时，kill 时触发。kill -9 时的信号 SIGKILL 不能捕捉，所以不用添加
	signal.Notify(quit, syscall.SIGTERM)

	// 等待退出信号
	sig := <-quit
	log.Printf("received signal: %v\n", sig)

	// 收到信号后，优雅地关闭服务器
	log.Println("server shutting down")
	fmt.Println("=============")
	zap.S().Named("demo").Info("程序退出===========", sig)
	if err = register_client.DeRegister(serviceId); err != nil {
		zap.S().Named("demo").Info("注销失败:", err.Error())
	} else {
		zap.S().Named("demo").Info("注销成功:")
	}
}
