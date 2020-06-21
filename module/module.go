package module

import (
	"github.com/ryus08/jiraTagger/config"
	"github.com/ryus08/jiraTagger/controller"
	"go.uber.org/dig"
)

type TaggerModule struct {
}

func (module *TaggerModule) Load(container *dig.Container) error {
	return container.Provide(config.Load)
}

func (module *TaggerModule) LoadRequestContext(appContainer *dig.Container, requestContainer *dig.Container) error {
	if err := appContainer.Invoke(func(c *config.Config) error {
		return requestContainer.Provide(func() *config.Config {
			return c
		})
	}); err != nil {
		return err
	}

	return requestContainer.Provide(func(config *config.Config) *controller.TaggerController {
		return &controller.TaggerController{Config: config}
	})
}
