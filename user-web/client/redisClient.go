package client

import (
	"github.com/go-redis/redis/v8"
)

//var rdb = *redis.NewClient(&redis.Options{
//	Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
//	Password: global.ServerConfig.RedisInfo.Password,
//	DB:       global.ServerConfig.RedisInfo.DB,
//})

var (
	Rdb *redis.Client
)

//func initClient() (err error) {
//	redis.NewClient(&redis.Options{
//		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
//		Password: global.ServerConfig.RedisInfo.Password,
//		DB:       global.ServerConfig.RedisInfo.DB,
//	})
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, err = rdb.Ping(ctx).Result()
//	return err
//}
