package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

type Config struct {
	SigningSecret string
}

type MyController struct {
}

func (myController *MyController) Handle(c *gin.Context) {
	receive := &controller.Receive{}
	var body string
	if c.Request.Method == "GET" {
		body = "{Content: \"Hello!\"}"
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		body = buf.String()
	}

	// TODO: Make this a 403
	e := receive.Authorize(c.Request.Header, &body)
	var response interface{}
	var statusCode int
	if e == nil {
		response, e = receive.Handler(&body)
		statusCode = http.StatusOK
	}

	if e != nil {
		response = e
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, response)
}

type ContextualHandler struct {
	GlobalContainer *dig.Container
	Module          func(*dig.Container)
}

type Controller interface {
	Handle(*gin.Context)
}

func test(controller Controller) {
}

func newMyController(config *Config) Controller {
	controller := &MyController{}
	test(controller)
	return controller
}

func (contextualHandler *ContextualHandler) getContextualHandler(controllerProvider interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestContextContainer := dig.New()
		requestContextContainer.Provide(func() *gin.Context {
			return c
		})
		// TOOD: Auto-wire up everything global onto request scoped.
		// Right now we're assuming the RequestScope Module pumps things through from the global container
		if contextualHandler.Module != nil {
			contextualHandler.Module(requestContextContainer)
		}
		requestContextContainer.Provide(controllerProvider)

		err := requestContextContainer.Invoke(func(handler Controller) {
			handler.Handle(c)
		})
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	container := dig.New()
	contextualHandler := &ContextualHandler{
		GlobalContainer: container,
		Module: func(rc *dig.Container) {
			container.Invoke(func(config *Config) {
				rc.Provide(func() *Config {
					return config
				})
			})
		},
	}

	container.Provide(func() *Config {
		var config Config
		viper.SetEnvPrefix("JIRATAGGER")
		viper.AutomaticEnv()
		viper.BindEnv("SIGNINGSECRET")

		err := viper.Unmarshal(&config)
		if err != nil {
			fmt.Println("unable to decode into struct, %v", err)
			panic(err)
		}
		return &config
	})

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(corsConfig))

	router.GET("/dev/receive", contextualHandler.getContextualHandler(newMyController))
	router.POST("/dev/receive", contextualHandler.getContextualHandler(newMyController))

	router.Run(":3000")
}
