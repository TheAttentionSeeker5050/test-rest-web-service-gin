package config

// import packages
import (
	//PostgreSQL Driver
	"errors"
	"fmt"
	"os"
	"workspace/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error

	// get the environment variables of the database and store them separately
	host := os.Getenv("DB_LOCAL_HOST")
	port := os.Getenv("DB_LOCAL_PORT")
	user := os.Getenv("DB_LOCAL_USER")
	password := os.Getenv("DB_LOCAL_PASSWORD")
	dbname := os.Getenv("DB_LOCAL_NAME")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "192.168.0.99", "nicolas", "mysecretpassword", "my_first_golang_db", "32831") // this is on docker container

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port) // this is on docker container

	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// auto migrate the models
	Database.AutoMigrate(&model.TestModel{})
	Database.AutoMigrate(&model.CalculatorHistoryModel{})

	if err != nil {
		// if there is an error opening the connection, handle it
		return nil, errors.New("Error on DB connection:\n" + err.Error())
	} else {
		// if the connection was successful, return nil error
		fmt.Println("Successfully connected to the database")
		return Database, nil
	}
}
