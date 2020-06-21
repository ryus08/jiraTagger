package main

import (
	"bytes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
	"github.com/ryus08/jiraTagger/digr"
	"github.com/ryus08/jiraTagger/module"
)

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

	e := receive.Authorize(c.Request.Header, &body)
	var response interface{}
	var statusCode int
	if e != nil {
		response = e
		statusCode = http.StatusUnauthorized
	} else {
		response, e = receive.Handler(&body)
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, response)
}

func main() {
	contextualHandler := digr.NewContextualHandler(&module.TaggerModule{})

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(corsConfig))

	router.GET("/dev/receive", contextualHandler.GetContextualHandler(
		func(config *module.Config) digr.Controller { return &MyController{} }))
	router.POST("/dev/receive", contextualHandler.GetContextualHandler(
		func(config *module.Config) digr.Controller { return &MyController{} }))

	router.Run(":3000")
}
