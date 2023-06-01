package config

// import packages
import (
	//PostgreSQL Driver
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var Database *gorm.DB

func ConnectDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "127.0.0.1", "postgres", "mysecretpassword", "test_db", "5432")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// if there is an error opening the connection, handle it
		return errors.New("Error on DB connection:\n" + err.Error())
	} else {
		// if the connection was successful, return nil error
		fmt.Println("Successfully connected to the database")
		return nil
	}
}
