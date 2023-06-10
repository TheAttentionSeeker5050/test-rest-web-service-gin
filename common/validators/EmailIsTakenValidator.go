package validators

import (
	"workspace/model"

	"gorm.io/gorm"
)

// add password validation method here, it should return true if the username is taken
func ValidateEmailIsTaken(db *gorm.DB, email string) bool {
	// check if the user name is already taken

	// create a user model instance
	var userModelInstance model.UserModel

	// get the user model instance by user name
	result := model.GetUserModelInstanceByEmail(db, &userModelInstance, email)

	// if email was found in the db, return true
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}

}
