package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Provider interface {
	GetString(key string) string
}

type Config struct {
	*viper.Viper
}

var (
	_             Provider = (*Config)(nil)
	once          sync.Once
	defaultConfig *Config
	err           error
	loadConfig    = defaultLoadConfig
)

func defaultLoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	v.SetDefault("port", 8080)
	v.SetDefault("log_level", "info")
	v.AutomaticEnv()

	err = v.ReadInConfig()
	return &Config{v}, err
}

// New returns the application configuration.
func New() (*Config, error) {
	once.Do(func() {
		defaultConfig, err = loadConfig()
	})
	return defaultConfig, err
}
