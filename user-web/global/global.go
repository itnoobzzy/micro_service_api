package global

import (
	ut "github.com/go-playground/universal-translator"

	"micro/user-web/config"
	"micro/user-web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	NacosConfig   *config.NacosConfig  = &config.NacosConfig{}
	UserSrvClient proto.UserClient
)
