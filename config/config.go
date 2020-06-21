package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	SigningSecret string
}

func Load() (*Config, error) {
	var config Config
	viper.SetEnvPrefix("JIRATAGGER")
	viper.AutomaticEnv()
	viper.BindEnv("SIGNINGSECRET")

	err := viper.Unmarshal(&config)
	return &config, err
}
