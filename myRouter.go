package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func myRouter() *gin.Engine {
	routerObj := gin.Default()
	routerObj.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	routerObj.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})

	routerObj.Run(":5000")

	return routerObj
}
