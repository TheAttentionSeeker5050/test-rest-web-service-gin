package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ory/dockertest/v3"
	_ "github.com/ory/dockertest/v3/docker"
)

func MockDBSetup() (*gorm.DB, error) {
	// Initialize the Docker pool
	pool, _ := dockertest.NewPool("")

	// Run a PostgreSQL container
	resource, _ := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=secret"})
	var database *gorm.DB
	// Wait for the container to be ready
	err := pool.Retry(func() error {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=192.168.0.99 port=32783 user=postgres password=secret dbname=postgres sslmode=disable"),
			// DSN: fmt.Sprintf("host=localhost port=%s user=postgres password=secret dbname=postgres sslmode=disable", resource.GetPort("5432/tcp")),
		}), &gorm.Config{})
		if err != nil {
			return err
		} else {
			database = db
			return nil
		}
	})

	resource = resource

	return database, err
}
