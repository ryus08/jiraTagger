package main

import (
	"bytes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
)

func receive(c *gin.Context) {
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

func main() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost"}
	router.Use(cors.New(corsConfig))

	router.GET("/dev/receive", receive)
	router.POST("/dev/receive", receive)

	router.Run(":3000")
}
