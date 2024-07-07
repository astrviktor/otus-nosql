package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

const ServicePrefix = "service"

type Config struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	Host     string `default:"127.0.0.1"`
	Port     string `default:"8081"`
	DataSize int    `default:"20480"` // 20Mb

	Redis RedisConfig
}

type RedisConfig struct {
	Host     string `default:"127.0.0.1"`
	Port     string `default:"6379"`
	Password string `default:"password"`
	DB       int    `default:"0"`
}

func ReadConfig(prefix string) (*Config, error) {
	cfg := Config{}

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, fmt.Errorf("fail to parse config: %s\n", err.Error())
	}

	return &cfg, nil
}

func PrintUsage(prefix string) {
	cfg := Config{}
	err := envconfig.Usage(prefix, &cfg)
	if err != nil {
		fmt.Printf("fail to print envconfig usage: %s", err.Error())
	}
}
