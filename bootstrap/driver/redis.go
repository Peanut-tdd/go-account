package driver

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	R_DB = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     GVA_VP.GetString("redis.address"),  // string
		Password: GVA_VP.GetString("redis.password"), // string
		DB:       GVA_VP.GetInt("redis.database"),    // int
	})

	RedisPong, RedisErr := R_DB.Ping(context.Background()).Result() // 心跳

	if RedisErr != nil {
		log.Println("Redis服务未运行。。。", RedisPong, RedisErr)
		log.Println("Redis常用命令：\n" +
			" 启动：src/redis-server \n" +
			" 进入命令行：src/redis-cli \n" +
			" 关闭安全模式：CONFIG SET protected-mode no \n" +
			" 重置密码：config set requirepass [密码]\n")
		//os.Exit(200)
	} else {
		log.Println("GoRedis已连接 >>> ")
	}
}
