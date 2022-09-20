package utils

import (
	"account_check/bootstrap/driver"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx context.Context

func initRedis() context.Context {

	return context.Background()
}

// 存普通string类型，10分钟过期,time.Minute*10
func RedisSet(key string, value interface{}, expire time.Duration) {
	ctx := initRedis()
	err := driver.R_DB.Set(ctx, key, value, expire).Err()
	if err != nil {
		fmt.Printf("set key failed, err:%v\n", err)
	}
	return

}

func RedisGet(key string) interface{} {
	ctx := initRedis()
	value, err := driver.R_DB.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
	}

	return value
}
