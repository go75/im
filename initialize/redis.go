package initialize

import (
	"im/global"
	"im/log"

	"github.com/go-redis/redis"
)

func initRedis() {
	global.Rd = redis.NewClient(&redis.Options {
		Addr: global.Config.Redis.Addr,
		DB: global.Config.Redis.DB,
		PoolSize: global.Config.Redis.PoolSize,
		MinIdleConns: global.Config.Redis.MinIdleConns,
	})
	pong, err := global.Rd.Ping().Result()
	if err != nil {
		panic(err)
	} else {
		log.Info.Println("redis inited, ", pong)
	}
}