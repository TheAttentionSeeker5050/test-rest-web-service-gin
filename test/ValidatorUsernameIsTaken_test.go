package test

import (
	"testing"
	"workspace/common/validators"
	"workspace/config"
	"workspace/model"

	"github.com/stretchr/testify/assert"
)

// we are going to test the validator for username already taken
func TestValidateUsernameIsTaken(t *testing.T) {

	// first, we need to create a database connection using the MockDBSetup() function
	db, err := config.MockDBSetup(t)
	assert.Nil(t, err)

	// migrate the user model
	db.AutoMigrate(&model.UserModel{})

	// drop all previous entries from the table
	db.Where("id > ?", 0).Delete(&model.UserModel{})

	// we will create a user model instance and save it to the database just once
	// then we will use the same user model instance to test the validator

	// create a user model instance
	userModelInstance := model.UserModel{
		UserName: "test5555",
		Email:    "test5555@gmail.com",
		PassWord: "password",
	}

	// create the user model instance in the database
	result := model.CreateUserModelInstance(db, &userModelInstance)

	// check if there is an error
	assert.Nil(t, result.Error)

	// check if the user model instance was created
	assert.Equal(t, int64(1), result.RowsAffected)

	// now we will test the validator
	// test case 1: username is not taken
	// we will use the same user model instance that we created above
	// we will pass a random username to the validator
	// the validator should return false
	usernameIsTaken := validators.ValidateUserNameIsTaken(db, "thisisanewusername")

	// check if the username is taken
	assert.False(t, usernameIsTaken)

	// test case 2: username is taken
	// we will use the same user model instance that we created above
	// we will pass the username of the user model instance to the validator
	// the validator should return true
	usernameIsTaken = validators.ValidateUserNameIsTaken(db, userModelInstance.UserName)

	// check if the username is taken
	assert.True(t, usernameIsTaken)

}
