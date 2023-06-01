package config

// import packages
import (
	//PostgreSQL Driver
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var Database *gorm.DB

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "127.0.0.1", "postgres", "mysecretpassword", "test_db", "5432")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error on DB connection")
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
