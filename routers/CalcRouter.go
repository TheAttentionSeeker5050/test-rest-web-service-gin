package routers

import (
	"workspace/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// we should also import the database as we are going to use it in some of our routes
func CalcRouter(routerObj *gin.Engine, db *gorm.DB) *gin.Engine {

	// add the address prefix
	var addressPrefix string = "/api/v1/calculators"

	// add the router group using the address prefix
	calculators := routerObj.Group(addressPrefix)

	// calculator routes
	calculators.POST("/basic-calc", func(ctx *gin.Context) {
		controller.BasicCalcController(ctx, db)
	})
	calculators.POST("/bin-to-hex", func(ctx *gin.Context) {
		controller.BinToHexController(ctx, db)
	})
	calculators.POST("/hex-to-bin", func(ctx *gin.Context) {
		controller.HexToBinController(ctx, db)
	})
	calculators.POST("/statistics-calc", func(ctx *gin.Context) {
		controller.StatisticsCalculatorController(ctx, db)
	})

	// calculator history routes
	calculators.POST("/history", func(ctx *gin.Context) {
		controller.CalcHistorySaveToDbController(ctx, db)
	}) // this is just for testing purposes
	calculators.GET("/history", func(ctx *gin.Context) {
		controller.CalcHistoryGetAllController(ctx, db)
	})

	return routerObj
}
