package test

import (
	"testing"
	"workspace/config"
	"workspace/model"
)

func TestConnectDB(t *testing.T) {
	_, err := config.ConnectDB()

	if err != nil {
		t.Errorf("There was an error with the db connection: %v ", err)
	} else {
		t.Log("Success")
	}

}

func TestSampleQuery(t *testing.T) {
	// the struct that will hold the result
	var testModel model.TestModel

	// the struct that will hold the data to be saved
	testModel = model.TestModel{
		UserId:   1,
		Email:    "email@email.com",
		Password: "password",
		UserName: "username",
	}

	// connect to the database
	db, err := config.ConnectDB()

	// check if there is an error with the connection
	if err != nil {
		t.Errorf("There was an error with the db connection: %v ", err)
	} else {
		t.Log("Success")
	}

	db.AutoMigrate(&model.TestModel{})

	// create a save query
	result := db.Create(&testModel)

	// check if there is an error with the query
	if result.Error != nil {
		t.Errorf("There was an error with the create query: %v", result.Error)
	} else {
		t.Log("Success")
	}

	// check the result of the query
	if result.RowsAffected != 1 {
		t.Errorf("Expected %v but got %v", 1, result.RowsAffected)
	} else {
		t.Log("Success")
	}

	// the struct that will hold the result
	var resultStruct model.TestModel

	// create a find query
	db.First(&resultStruct)

	// check if the result is correct
	if resultStruct.UserId != 1 {
		t.Errorf("Expected %v but got %v", 1, resultStruct.UserId)
	} else {
		t.Log("Success")
	}

}
