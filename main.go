package main

import (
	"fmt"
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

	// configure Oauth2
	oauth2ConfigObj := config.Oauth2Config()

	fmt.Println("auth config:", oauth2ConfigObj)

	// Connect to the database
	database := loadDatabase()

	// declare router and routes
	router := gin.Default()
	router.NoRoute(controller.NoRouteOrMethodController)
	router = routers.CalcRouter(router, database) // calculator routes

	// run router on port 5000
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
