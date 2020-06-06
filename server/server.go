package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
)

func receive(c *gin.Context) {
	receive := &controller.Receive{}

	var requestBody controller.RequestBody

	if c.Request.Method == "GET" {
		requestBody = controller.RequestBody{Content: "Hello!"}
	} else {
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	response := receive.Handler(&requestBody)
	c.JSON(200, response)
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
