package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"todolist/conf"
)

var RedisClient *redis.Client

func RedisInit() {
	rConfig := conf.Config.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.Host, rConfig.Port),
		Password: rConfig.Password,
		DB:       rConfig.Db,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	RedisClient = client
	fmt.Println("Redis initialized")
}
