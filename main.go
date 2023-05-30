package main

import (
	"workspace/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router = SampleRouter(router)
	router = routers.CalcRouter(router)
	router.Run(":5000")

}
