package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

func SampleRouter(routerObj *gin.Engine) *gin.Engine {
	// routerObj := gin.Default()
	routerObj.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// routerObj.GET("/os", func(c *gin.Context) {
	// 	c.JSON(200, runtime.GOOS)
	// })

	routerObj.GET("/os", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"OS": runtime.GOOS,
		})
	})

	// routerObj.Run(":5000")

	return routerObj
}
