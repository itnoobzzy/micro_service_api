package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"micro/demo/nacos_test/config"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "192.168.146.189",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "da7c2a11-8ae4-44d6-b184-97adf1b2b6a6", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	serverConfig := config.ServerConfig{}
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)
	//err = configClient.ListenConfig(vo.ConfigParam{
	//	DataId: "user-web.json",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件变化")
	//		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	//	},
	//})
	//time.Sleep(3000 * time.Second)

}
