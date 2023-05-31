package main

import (
	"workspace/controller"
	"workspace/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router = SampleRouter(router)
	router.NoRoute(controller.NoRouteController)

	router = routers.CalcRouter(router)
	router.Run(":5000")

}
