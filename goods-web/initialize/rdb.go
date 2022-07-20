package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	rdbClient "micro/user-web/client"
	"micro/user-web/global"
	"time"
)

func InitRdbClient() (err error) {
	rdbClient.Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password,
		DB:       global.ServerConfig.RedisInfo.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdbClient.Rdb.Ping(ctx).Result()
	return err
}
