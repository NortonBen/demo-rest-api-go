package redis

import (
	"apm/micro/load"
	base "github.com/go-redis/redis"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/logger"
	"github.com/urfave/cli/v2"
)

type Redis struct {
	client     *base.Client
	Config     *RedisConfig
	configName string
}

func (r *Redis) Flag() []cli.Flag {
	return []cli.Flag{}
}

func (r *Redis) Start() error {

	var err error
	val, err := config.Get(r.configName)
	if err != nil {
		logger.Error("Config Get", err)
		return err
	}
	err = val.Scan(&r.Config)
	if err != nil {
		logger.Error("Config Get", err)
		return err
	}

	r.client, err = ConnectRedis(r.Config)

	return err
}

func (r *Redis) Loader() load.Loader {
	return r
}

func NewRedis(configName string) *Redis {
	return &Redis{
		Config:     &RedisConfig{},
		configName: configName,
	}
}

func (r *Redis) Client() *base.Client {
	return r.client
}

func (r Redis) Name() string {
	return "redis"
}

func (r *Redis) Stop() error {
	return r.client.Close()
}
