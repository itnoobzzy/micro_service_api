package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"micro/user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	zap.S().Named("demo").Infof("DEBUG: &v", viper.Get(env))
	return viper.GetBool(env)
}

func getConfig(cc config_client.IConfigClient) {
	content, err := cc.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Named("demo").Fatalf("读取nacos配置失败： %s", err.Error())
		panic(err)
	}
	zap.S().Named("demo").Info(&global.ServerConfig)
}

func InitConfig() {
	debug := GetEnvInfo("DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Named("demo").Infof("配置信息: &v", global.NacosConfig)
	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize: 1,
			MaxAge:  1,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	//content, err := configClient.GetConfig(vo.ConfigParam{
	//	DataId: global.NacosConfig.DataId,
	//	Group:  global.NacosConfig.Group})
	//
	//if err != nil {
	//	panic(err)
	//}
	getConfig(configClient)
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			getConfig(configClient)
		},
	})

	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	//err = json.Unmarshal([]byte(content), &global.ServerConfig)
	//if err != nil {
	//	zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	//}
	//zap.S().Info(&global.ServerConfig)

}
