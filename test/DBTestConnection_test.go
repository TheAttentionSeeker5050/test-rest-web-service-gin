package test

import (
	"testing"
	"workspace/config"
	"workspace/model"

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
	// initialize the database
	db, err := config.MockDBSetup(t)

	// check if there is an error with the db connection
	assert.Nil(t, err)

	// define the struct that will hold the result
	var testModel model.TestModel

	// the struct that will hold the data to be saved
	testModel = model.TestModel{
		UserId:   1,
		Email:    "email@email.com",
		Password: "password",
		UserName: "username",
	}

	db.AutoMigrate(&model.TestModel{})

	// drop all previous entries from the table
	db.Where("user_id > ?", 0).Delete(&model.TestModel{})

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

	// drop all entries from the table
	db.Where("user_id > ?", 0).Delete(&model.TestModel{})

}
