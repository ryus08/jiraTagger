package main

import (
	"net/http"

	"time"

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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "localhost"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/dev/receive", receive)
	router.POST("/dev/receive", receive)

	router.Run(":3000")
}
