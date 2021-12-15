package config

import (
	"github.com/cdo-pand/simple-rest-api-with-golang/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Config singleton

type Config struct {
	IsDebug *bool `yaml:"is_debug" env:"IS_DEBUG" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env:"TYPE" env-default:"port"`
		BindIP string `json:"bind_ip" env:"BIND_IP" env-default:"127.0.0.1"`
		Port   string `json:"port" env:"PORT" env-default:"8080"`
	} `yaml:"listen"`
	MongoDB struct {
		Host       string `json:"host"`
		Port       string `json:"port"`
		Database   string `json:"database"`
		AuthDB     string `json:"auth_db"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Collection string `json:"collection"`
	} `json:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
