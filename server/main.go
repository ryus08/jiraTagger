package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryus08/jiraTagger/controller"
)

func receive(c *gin.Context) {
	receive := &controller.Receive{}

	var requestBody controller.RequestBody
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := receive.Handler(&requestBody)
	c.JSON(200, response)
}

func main() {
	router := gin.Default()
	router.GET("/dev/receive", receive)
	router.POST("/dev/receive", receive)

	router.Run(":3000")
}
