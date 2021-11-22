package redis

import (
	base "github.com/go-redis/redis"
	"github.com/micro/micro/v3/service/logger"
)

func ConnectRedis(redisConfig *RedisConfig) (*base.Client, error) {

	logger.Debug("config Redis", redisConfig)

	client := base.NewClient(&base.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		logger.Errorf("Ping Redis: ", err.Error())
	}
	return client, err
}
