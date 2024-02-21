package tool

import "github.com/go-redis/redis/v8"

type RedisEngine struct {
	*redis.Client
}

var Redis RedisEngine

func InitRedis(config *Config) *redis.Client {
	redisConfig := config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	Redis.Client = rdb
	return rdb

}
