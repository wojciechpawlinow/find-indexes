package config

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func mockLoadConfigSuccess() (*config, error) {
	mockViper := viper.New()
	mockViper.Set("port", 8080)
	mockViper.Set("log_level", "debug")

	return &config{mockViper}, nil
}

func TestNewConfig(t *testing.T) {
	loadConfig = mockLoadConfigSuccess
	defer func() { loadConfig = defaultLoadConfig }()

	cfg, err := New()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, 8080, cfg.GetInt("port"))
	assert.Equal(t, "debug", cfg.GetString("log_level"))
}

// no more tests scenarios due to once.Do() usage
