package routers

import (
	"workspace/controller"

	"github.com/gin-gonic/gin"
)

func CalcRouter(routerObj *gin.Engine) *gin.Engine {

	// add the address prefix
	var addressPrefix string = "/api/v1/calculators"

	// add the router group using the address prefix
	calculators := routerObj.Group(addressPrefix)

	// the real requests:

	calculators.POST("/basic-calc", controller.BasicCalcController)
	calculators.POST("/bin-to-hex", controller.BinToHexController)
	calculators.POST("/hex-to-bin", controller.HexToBinController)
	calculators.POST("/statistics-calc", controller.StatisticsCalculatorController)
	calculators.GET("/history", func(c *gin.Context) {
		c.String(200, "**API Calculator History** \nThis should make use of the server database")
	})

	return routerObj
}
