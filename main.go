package main

import (
	"workspace/config"
	"workspace/controller"
	"workspace/routers"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
)

var ORM orm.Ormer

func main() {

	// Connect to the database
	loadDatabase()

	// config.ConnectToDb()
	// ORM = config.GetOrmObject()

	router := gin.Default()
	router.NoRoute(controller.NoRouteOrMethodController)

	router = routers.CalcRouter(router)
	router.Run(":5000")

}

func loadDatabase() {
	err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
}
