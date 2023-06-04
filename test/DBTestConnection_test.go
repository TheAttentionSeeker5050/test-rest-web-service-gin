package test

import (
	"fmt"
	"testing"
	"workspace/config"
	"workspace/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ory/dockertest/v3"
	_ "github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	_, err := config.ConnectDB()

	if err != nil {
		t.Errorf("There was an error with the db connection: %v ", err)
	} else {
		t.Log("Success")
	}

}

func TestSampleQueryUsingMockDB(t *testing.T) {
	// Initialize the Docker pool
	pool, err := dockertest.NewPool("")
	assert.NoError(t, err)

	// Run a PostgreSQL container
	resource, err := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=secret"})
	assert.NoError(t, err)

	// Wait for the container to be ready
	err = pool.Retry(func() error {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=192.168.0.99 port=32783 user=postgres password=secret dbname=postgres sslmode=disable"),
			// DSN: fmt.Sprintf("host=localhost port=%s user=postgres password=secret dbname=postgres sslmode=disable", resource.GetPort("5432/tcp")),
		}), &gorm.Config{})
		assert.Nil(t, err)

		// Run your database migration or setup logic here

		// define the struct that will hold the result
		var testModel model.TestModel

		// the struct that will hold the data to be saved
		testModel = model.TestModel{
			UserId:   1,
			Email:    "email@email.com",
			Password: "password",
			UserName: "username",
		}

		// // connect to the database
		// db, err := config.ConnectDB()

		// auto migrate any changes
		db.AutoMigrate(&model.TestModel{})

		// create a save query
		result := model.CreateTestModelInstance(db, &testModel)

		// check if there is an error with the query
		assert.Nil(t, result.Error)

		// check if the data was saved by checking the result of the query
		assert.Equal(t, 1, int(result.RowsAffected))

		// the struct that will hold the result
		var resultStruct model.TestModel

		// create a find query
		model.GetLastTestModelInstance(db, &resultStruct)

		// check if the result is correct
		assert.Equal(t, 1, resultStruct.UserId)
		assert.Equal(t, "email@email.com", resultStruct.Email)
		assert.Equal(t, "password", resultStruct.Password)
		assert.Equal(t, "username", resultStruct.UserName)

		return nil
	})
	assert.NoError(t, err)

	// Perform your test operations using the mock PostgreSQL database
	resource.GetPort("5432/tcp") // returns the exposed port for the PostgreSQL instance
	// When you're done, kill and remove the container and check if there is any error
	assert.NoError(t, pool.Purge(resource))

}
