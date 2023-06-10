package routers

import (
	"workspace/controller/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// here we will add the router group for the user routes
func UserRouter(routerObj *gin.Engine, db *gorm.DB) *gin.Engine {
	// add the address prefix for the user routes
	var addressPrefix string = "/api/v1/users"

	// add the router group using the address prefix
	userRoutes := routerObj.Group(addressPrefix)

	// add the user routes
	userRoutes.POST("/register", func(ctx *gin.Context) {
		// call the controller
		user.RegisterUserController(ctx, db)
	})

	userRoutes.POST("/login", func(ctx *gin.Context) {
		// display text for testing purposes
		user.LoginUserController(ctx, db)
	})

	userRoutes.POST("/logout", func(ctx *gin.Context) {
		// display text for testing purposes
		ctx.JSON(200, gin.H{
			"message": "logout",
		})
	})

	userRoutes.POST("/update", func(ctx *gin.Context) {
		// display text for testing purposes
		ctx.JSON(200, gin.H{
			"message": "update",
		})
	})

	userRoutes.POST("/delete", func(ctx *gin.Context) {
		// display text for testing purposes
		ctx.JSON(200, gin.H{
			"message": "delete",
		})
	})

	// wiew user profile
	userRoutes.GET("/profile", func(ctx *gin.Context) {
		// display text for testing purposes
		user.GetUserProfileDataController(ctx, db)
	})

	return routerObj
}
