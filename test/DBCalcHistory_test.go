package test

import (
	"fmt"
	"testing"
	"workspace/config"
	"workspace/model"

	"github.com/ory/dockertest/v3"
	_ "github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// here we will test that the user can save and retrieve the calc history to and from the db
func TestCalcHistorySaveToDb(t *testing.T) {
	var calcHistoryModel model.CalculatorHistoryModel

	// the struct that will hold the data to be saved
	calcHistoryModel = model.CalculatorHistoryModel{
		UserName:       "anonymous",
		CalculatorType: "BasicCalculator",
		Params:         "1,2,'+'",
		Results:        "3",
	}

	// connect to the database
	db, err := config.MockDBSetup()

	// check if there is an error with the connection
	if err != nil {
		t.Errorf("There was an error with the db connection: %v ", err)
	} else {
		t.Log("Success")
	}

	// auto migrate to update our model if needed
	db.AutoMigrate(&model.CalculatorHistoryModel{})

	// create a save query
	result := db.Create(&calcHistoryModel)

	// check if there is an error with the query
	if result.Error != nil {
		t.Errorf("There was an error with the create query: %v", result.Error)
	} else {
		t.Log("Success")
	}

	// validate if the data was saved on the db using the result of the query
	if result.RowsAffected != 1 {
		t.Errorf("Expected %v but got %v", 1, result.RowsAffected)
	} else {
		t.Log("Success")
	}

	// create a query to retrieve the data from the db
	var calcHistoryModelFromDb model.CalculatorHistoryModel

	// create a query to retrieve the data from the db and save it on the calcHistoryModelFromDb structure
	db.Last(&calcHistoryModelFromDb)

	// validate if the data retrieved from the db matches
	if calcHistoryModelFromDb.UserName != calcHistoryModel.UserName {
		t.Errorf("Expected %v but got %v", calcHistoryModel.UserName, calcHistoryModelFromDb.UserName)
	} else {
		t.Log("Success")
	}

	if calcHistoryModelFromDb.CalculatorType != calcHistoryModel.CalculatorType {
		t.Errorf("Expected %v but got %v", calcHistoryModel.CalculatorType, calcHistoryModelFromDb.CalculatorType)
	} else {
		t.Log("Success")
	}

	if calcHistoryModelFromDb.Params != calcHistoryModel.Params {
		t.Errorf("Expected %v but got %v", calcHistoryModel.Params, calcHistoryModelFromDb.Params)
	} else {
		t.Log("Success")
	}

	if calcHistoryModelFromDb.Results != calcHistoryModel.Results {
		t.Errorf("Expected %v but got %v", calcHistoryModel.Results, calcHistoryModelFromDb.Results)
	} else {
		t.Log("Success")
	}

}

// here we will test that the user can remove the calc history from the db
func TestCalcHistoryDBQueries(t *testing.T) {
	// Initialize the Docker pool
	pool, err := dockertest.NewPool("")
	assert.NoError(t, err)

	// Run a PostgreSQL container
	resource, err := pool.Run("postgres", "13", []string{"POSTGRES_PASSWORD=secret"})
	assert.NoError(t, err)

	// Wait for the container to be ready
	err = pool.Retry(func() error {
		db, err := gorm.Open(postgres.New(postgres.Config{
			// DSN: fmt.Sprintf("host=localhost port=%s user=postgres password=secret dbname=postgres sslmode=disable", resource.GetPort("5432/tcp")),
			DSN: fmt.Sprintf("host=192.168.0.99 port=32783 user=postgres password=secret dbname=postgres sslmode=disable"),
		}), &gorm.Config{})
		assert.Nil(t, err)

		// Run your database migration or setup logic here

		// define the struct that will hold the result
		var calcHistoryModel model.CalculatorHistoryModel

		// the struct that will hold the data to be saved
		calcHistoryModel = model.CalculatorHistoryModel{
			UserName:       "anonymous",
			CalculatorType: "BasicCalculator",
			Params:         "1,2,'+'",
			Results:        "3",
		}

		// auto migrate any changes
		db.AutoMigrate(&model.CalculatorHistoryModel{})

		// create a save query
		resultCreate := model.CreateCalculatorHistoryModelInstance(db, &calcHistoryModel)

		// check if there is an error with the query
		assert.Nil(t, resultCreate.Error)

		// validate if the data was saved on the db using the result of the query
		assert.Equal(t, 1, int(resultCreate.RowsAffected))

		// create a query to retrieve the data from the db
		var calcHistoryModelFromDb model.CalculatorHistoryModel

		// create a query to retrieve the data from the db and save it on the calcHistoryModelFromDb structure
		model.GetLastCalculatorHistoryModelInstance(db, &calcHistoryModelFromDb)

		// validate if the data retrieved from the db matches
		assert.Equal(t, calcHistoryModel.UserName, calcHistoryModelFromDb.UserName)
		assert.Equal(t, calcHistoryModel.CalculatorType, calcHistoryModelFromDb.CalculatorType)
		assert.Equal(t, calcHistoryModel.Params, calcHistoryModelFromDb.Params)
		assert.Equal(t, calcHistoryModel.Results, calcHistoryModelFromDb.Results)

		// test the list query
		var calcHistoryModelList []model.CalculatorHistoryModel

		// create a query to retrieve the data from the db and save it on the calcHistoryModelFromDb structure
		model.GetAllCalculatorHistoryModelInstances(db, &calcHistoryModelList)

		// validate if the data retrieved from the db matches
		assert.Equal(t, calcHistoryModel.UserName, calcHistoryModelList[0].UserName)
		assert.Equal(t, calcHistoryModel.CalculatorType, calcHistoryModelList[0].CalculatorType)
		assert.Equal(t, calcHistoryModel.Params, calcHistoryModelList[0].Params)
		assert.Equal(t, calcHistoryModel.Results, calcHistoryModelList[0].Results)

		// validate if the size of the list is 1
		assert.Equal(t, 1, len(calcHistoryModelList))

		// drop all entries from the table
		db.Where("id > 0").Delete(&model.CalculatorHistoryModel{})

		return nil
	})
	assert.NoError(t, err)

	// Perform your test operations using the mock PostgreSQL database
	resource.GetPort("5432/tcp") // returns the exposed port for the PostgreSQL instance
	// When you're done, kill and remove the container and check if there is any error
	assert.NoError(t, pool.Purge(resource))
}
