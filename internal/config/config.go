package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Provider interface {
	GetString(key string) string
}

type config struct {
	*viper.Viper
}

var (
	once          sync.Once
	defaultConfig *config
	err           error
	loadConfig    = defaultLoadConfig
)

func defaultLoadConfig() (*config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	v.SetDefault("port", 8080)
	v.SetDefault("log_level", "info")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	return &config{v}, err
}

func New() (*config, error) {
	once.Do(func() {
		defaultConfig, err = loadConfig()
	})
	return defaultConfig, err
}
