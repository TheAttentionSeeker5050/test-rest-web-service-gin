package routers

import (
	"github.com/gin-gonic/gin"
)

func CalcRouter(routerObj *gin.Engine) *gin.Engine {

	var addressPrefix string = "/api/v1/calculators"

	routerObj.GET(addressPrefix+"/calc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello calc!",
		})
	})

	// routerObj.GET("/os", func(c *gin.Context) {
	// 	c.JSON(200, runtime.GOOS)
	// })

	routerObj.GET(addressPrefix+"/calc2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello calc2!",
		})
	})

	return routerObj
}
