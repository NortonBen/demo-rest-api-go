package util

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

func GetCacheRedis(client *redis.Client, name string, v interface{}) (error, bool) {
	byteData, err := client.Get(name).Bytes()
	if err == redis.Nil {
		return nil, false
	} else if err != nil {
		return err, false
	}
	err = json.Unmarshal(byteData, v)
	if err != nil {
		return err, false
	}
	return nil, true
}

func SetCacheRedis(client *redis.Client, name string, v interface{}, duration time.Duration) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return client.Set(name, jsonData, duration).Err()
}

func RemoveCacheRedis(client *redis.Client, name string) error {
	return client.Del(name).Err()
}
