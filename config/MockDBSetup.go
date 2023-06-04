package config

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ory/dockertest/v3"
	_ "github.com/ory/dockertest/v3/docker"
)

func MockDBSetup(t *testing.T) (*gorm.DB, error) {
	// Initialize the Docker pool
	// Create a new instance of the Docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Failed to create Docker pool: %s", err)
	}

	// Set the environment variables for the PostgreSQL container
	env := []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=mySecretPassword",
		"POSTGRES_DB=test_db",
		"DOCKER_MACHINE_NAME=postgres_container",
	}

	// Define the Docker container run options
	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env:        env,
	}

	// Start the PostgreSQL container
	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		t.Fatalf("Failed to start PostgreSQL container: %s", err)
	}
	defer func() {
		// Clean up the Docker container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Failed to clean up PostgreSQL container: %s", err)
		}
	}()

	// Get the host and port for connecting to the PostgreSQL container
	host := "192.168.0.99"
	// port := resource.GetPort("5433/tcp")
	port := "32783"

	// Create the DSN string for the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s port=%s user=postgres password=mySecretPassword dbname=test_db sslmode=disable", host, port)

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}

	return db, err
}
