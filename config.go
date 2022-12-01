package gobe

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	SqlConfig   SqlBaseConfig   `mapstructure:"sql" json:"sql"`
	MongoConfig MongoBaseConfig `mapstructure:"mongo" json:"mongo"`
	RedisConfig RedisBaseConfig `mapstructure:"redis" json:"redis"`
	// SwaggerConfig SwaggerBaseConfig `mapstructure:"swagger" json:"swagger"`
}

// Get application configuration from common configuration file (e.g JSON, YAML, etc.)
func GetConfigFromFile(filepath string) *Config {
	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	var conf *Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Fatal error marshal config file: %s \n", err)
	}
	return conf
}
