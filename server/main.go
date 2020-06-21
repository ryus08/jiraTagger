package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
	"github.com/ryus08/jiraTagger/digr"
	"github.com/ryus08/jiraTagger/module"
)

type ControllerWrapper struct {
	TaggerController *controller.TaggerController
}

// TODO: probably push the gin stuff right into the contextual handler
func (controllerWrapper *ControllerWrapper) Handle(c *gin.Context) {
	statusCode, response := controllerWrapper.TaggerController.Handle((c.Request))

	c.JSON(statusCode, response)
}

func main() {
	contextualHandler := digr.NewContextualHandler(&module.TaggerModule{})

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(corsConfig))

	controllerBuilder := func(taggerController *controller.TaggerController) digr.Controller {
		return &ControllerWrapper{TaggerController: taggerController}
	}

	router.GET("/dev/receive", contextualHandler.GetContextualHandler(controllerBuilder))
	router.POST("/dev/receive", contextualHandler.GetContextualHandler(controllerBuilder))

	router.Run(":3000")
}
