package test

import (
	"testing"
	"workspace/config"
	"workspace/model"
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
	db, err := config.ConnectDB()

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
func TestCalcHistoryRemoveFromDb(t *testing.T) {}
