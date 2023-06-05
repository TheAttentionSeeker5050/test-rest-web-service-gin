package config

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/ory/dockertest/v3"
	_ "github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

func MockDBSetup(t *testing.T) (*gorm.DB, error) {
	// get the absolute path of the env file
	// envPath, err := filepath.Join(filepath.Abs("../"), ".env")

	// Initialize the Docker pool
	// Load environment variables from .env file, this is for the local server, on production, the env variables are set on the server
	err := godotenv.Load("../.env")
	assert.Nil(t, err)

	// Create a new instance of the Docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Failed to create Docker pool: %s", err)
	}

	// Set the environment variables for the PostgreSQL container
	env := []string{
		// "POSTGRES_USER=postgres",
		"POSTGRES_USER=" + os.Getenv("DB_TEST_USER"),
		"POSTGRES_PASSWORD=" + os.Getenv("DB_TEST_PASSWORD"),
		"POSTGRES_DB=" + os.Getenv("DB_TEST_NAME"),
		"DOCKER_MACHINE_NAME=" + os.Getenv("DB_CONTAINER_NAME"),
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
	host := os.Getenv("DB_TEST_HOST")
	port := os.Getenv("DB_TEST_PORT")
	user := os.Getenv("DB_TEST_USER")
	password := os.Getenv("DB_TEST_PASSWORD")
	dbname := os.Getenv("DB_TEST_NAME")

	// Create the DSN string for the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}

	return db, err
}
