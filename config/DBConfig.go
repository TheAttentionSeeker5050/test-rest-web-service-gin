package config

// import packages
import (
	//PostgreSQL Driver
	"errors"
	"fmt"
	"workspace/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "192.168.0.99", "nicolas", "mysecretpassword", "my_first_golang_db", "32831") // this is on docker container
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
