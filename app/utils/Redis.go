package utils

import (
	"account_check/bootstrap/driver"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx context.Context

func initRedis() context.Context {

	return context.Background()
}

// 存普通string类型，10分钟过期,time.Minute*10
func RedisSet(key string, value interface{}, expire time.Duration) bool {
	ctx := initRedis()
	err := driver.R_DB.Set(ctx, key, value, expire).Err()
	if err != nil {
		return false
	}
	return true

}

func RedisGet(key string) interface{} {
	ctx := initRedis()
	value, err := driver.R_DB.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return nil
	}

	return value
}
