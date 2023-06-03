package main

import (
	"workspace/config"
	"workspace/controller"
	"workspace/routers"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ORM orm.Ormer

func main() {

	// Connect to the database
	database := loadDatabase()

	// config.ConnectToDb()
	// ORM = config.GetOrmObject()

	router := gin.Default()
	router.NoRoute(controller.NoRouteOrMethodController)

	router = routers.CalcRouter(router, database)
	router.Run(":5000")

}

func loadDatabase() *gorm.DB {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
		return nil
	}

	return db
}
