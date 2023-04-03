package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	route.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Hello World",
		})
	})

	route.Run()
}