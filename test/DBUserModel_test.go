package test

import (
	"testing"
	"workspace/config"
	"workspace/model"

	_ "github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

// here we will test the user model query methods with the db
func TestUserModelDBQueries(t *testing.T) {
	// first we will declare the user model struct
	var userModel model.UserModel

	// add sample data to the user model struct
	userModel = model.UserModel{
		UserName: "bert404",
		Email:    "berto@email.com",
		PassWord: "123456",
	}

	// connect to the database
	db, err := config.MockDBSetup(t)

	// check if there is an error with the connection
	assert.Nil(t, err)

	// auto migrate to update our model if needed
	db.AutoMigrate(&model.UserModel{})

	// delete all the data from the user model table
	db.Delete(&model.UserModel{})

	// create a save query
	result := model.CreateUserModelInstance(db, &userModel)

	// check if there is an error with the query
	assert.Nil(t, result.Error)

	// validate if the data was saved on the db using the result of the query
	assert.Equal(t, 1, int(result.RowsAffected))

	// create a query to retrieve the data from the db
	var userModelFromDb model.UserModel

	// create a query to retrieve the data from the db and save it on the userModelFromDb structure
	db.Last(&userModelFromDb)

	// validate if the data retrieved from the db matches
	assert.Equal(t, userModel.UserName, userModelFromDb.UserName)
	assert.Equal(t, userModel.Email, userModelFromDb.Email)
	assert.Equal(t, userModel.PassWord, userModelFromDb.PassWord)

	// create a query to retrieve the data from the db
	var userModelFromDbById model.UserModel

	// create a query to retrieve the data from the db and save it on the userModelFromDb structure
	db.First(&userModelFromDbById, userModelFromDb.Id)

	// validate if the data retrieved from the db matches
	assert.Equal(t, userModel.UserName, userModelFromDbById.UserName)
	assert.Equal(t, userModel.Email, userModelFromDbById.Email)
	assert.Equal(t, userModel.PassWord, userModelFromDbById.PassWord)

	// create a query to retrieve the data from the db by user name
	var userModelFromDbByUserName model.UserModel

	// create a query to retrieve the data from the db and save it on the userModelFromDb structure
	db.Where("user_name = ?", userModel.UserName).First(&userModelFromDbByUserName)

	// validate if the data retrieved from the db matches
	assert.Equal(t, userModel.UserName, userModelFromDbByUserName.UserName)
	assert.Equal(t, userModel.Email, userModelFromDbByUserName.Email)
	assert.Equal(t, userModel.PassWord, userModelFromDbByUserName.PassWord)

}
