// Package digr provides gin request-scoping for dig
package digr

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Module defines a set of providers onto the correct contextual container
type Module interface {
	Load(appContainer *dig.Container) error
	LoadRequestContext(appContainer *dig.Container, requestContainer *dig.Container) error
}

// ContextualHandler
type ContextualHandler struct {
	AppContainer *dig.Container
	Module       Module
}

func NewContextualHandler(module Module) *ContextualHandler {
	container := dig.New()
	module.Load(container)
	return &ContextualHandler{
		AppContainer: container,
		Module:       module,
	}
}

type Controller interface {
	Handle(*gin.Context)
}

func (contextualHandler *ContextualHandler) GetContextualHandler(controllerProvider interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestContextContainer := dig.New()
		if err := requestContextContainer.Provide(func() *gin.Context { return c }); err != nil {
			panic(err)
		}

		// TOOD: Auto-wire up everything global onto request scoped.
		// Right now we're assuming the RequestScope Module pumps things through from the global container
		if err := contextualHandler.Module.LoadRequestContext(
			contextualHandler.AppContainer,
			requestContextContainer); err != nil {
			panic(err)
		}

		if err := requestContextContainer.Provide(controllerProvider); err != nil {
			panic(err)
		}

		if err := requestContextContainer.Invoke(func(handler Controller) { handler.Handle(c) }); err != nil {
			panic(err)
		}
	}
}
