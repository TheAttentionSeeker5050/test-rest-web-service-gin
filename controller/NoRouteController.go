package controller

import "github.com/gin-gonic/gin"

func NoRouteController(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Not found",
	})
}
