package redis

type RedisConfig struct {
	Host     string `json:"address"`
	Password string `json:"password"`
}
