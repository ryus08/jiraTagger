package module

import (
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type Config struct {
	SigningSecret string
}

type TaggerModule struct {
}

func (module *TaggerModule) Load(container *dig.Container) error {
	return container.Provide(func() (*Config, error) {
		var config Config
		viper.SetEnvPrefix("JIRATAGGER")
		viper.AutomaticEnv()
		viper.BindEnv("SIGNINGSECRET")

		err := viper.Unmarshal(&config)
		return &config, err
	})
}

func (module *TaggerModule) LoadRequestContext(appContainer *dig.Container, requestContainer *dig.Container) error {
	return appContainer.Invoke(func(config *Config) error {
		return requestContainer.Provide(func() *Config {
			return config
		})
	})
}
