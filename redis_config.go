package gobe

import "github.com/go-redis/redis/v8"

type RedisBaseConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}

type RedisClient struct {
	*redis.Client
}

// Initialize new Redis client
func NewRedisClient(baseConfig *RedisBaseConfig) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     baseConfig.Host + ":" + baseConfig.Port,
		Username: baseConfig.Username,
		Password: baseConfig.Password,
		DB:       baseConfig.DB,
	})
	return RedisClient{client}
}
