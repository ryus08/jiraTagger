package main

import (
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Message string
}

func receive(c *gin.Context) {
	response := &ResponseBody{Message: "receive called"}
	c.JSON(200, response)
}

func main() {
	router := gin.Default()
	router.GET("/dev/receive", receive)
	router.POST("/dev/receive", receive)

	router.Run(":8080")
}
