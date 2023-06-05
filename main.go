package main

import (
	"log"
	"workspace/config"
	"workspace/controller"
	"workspace/routers"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var ORM orm.Ormer

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
		// return nil
	}

	return db
}
