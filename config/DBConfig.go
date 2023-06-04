package config

// import packages
import (
	//PostgreSQL Driver
	"errors"
	"fmt"
	"workspace/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "127.0.0.1", "postgres", "mysecretpassword", "test_db", "5432")
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

func connectDBForTest() (*gorm.DB, error) {
	var err error

	// connect to sqlite test db for testing (called test_db)
	Database, err = gorm.Open(sqlite.Open("test_db"), &gorm.Config{})

	// check if there is an error with the connection
	if err != nil {
		return nil, errors.New("Error on DB connection:\n" + err.Error())
	} else {
		return Database, nil
	}

}
