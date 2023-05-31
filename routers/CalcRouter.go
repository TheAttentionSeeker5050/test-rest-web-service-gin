package routers

import (
	"workspace/controller"

	"github.com/gin-gonic/gin"
)

func CalcRouter(routerObj *gin.Engine) *gin.Engine {

	var addressPrefix string = "/api/v1/calculators"

	// some test requests:
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

	// the real requests:

	routerObj.POST(addressPrefix+"/basic-calc", controller.BasicCalcController)
	routerObj.POST(addressPrefix+"/bin-to-hex", controller.BinToHexController)
	routerObj.POST(addressPrefix+"/hex-to-bin", controller.HexToBinController)
	routerObj.POST(addressPrefix+"/statistics-calc", controller.StatisticsCalculatorController)

	return routerObj
}
