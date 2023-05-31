package main

import (
	"workspace/controller"
	"workspace/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.NoRoute(controller.NoRouteOrMethodController)

	router = routers.CalcRouter(router)
	router.Run(":5000")

}
