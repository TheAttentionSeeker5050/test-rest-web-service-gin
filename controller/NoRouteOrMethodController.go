package controller

import "github.com/gin-gonic/gin"

func NoRouteOrMethodController(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Invalid request - Method not Allowed or Route not found",
	})
}
