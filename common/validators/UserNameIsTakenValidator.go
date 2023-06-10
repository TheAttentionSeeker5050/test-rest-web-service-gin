package validators

import (
	"workspace/model"

	"gorm.io/gorm"
)

// add password validation method here, it should return true if the username is taken
func ValidateUserNameIsTaken(db *gorm.DB, userName string) bool {
	// check if the user name is already taken

	// create a user model instance
	var userModelInstance model.UserModel

	// get the user model instance by user name
	result := model.GetUserModelInstanceByUserName(db, &userModelInstance, userName)

	// if username was found in the db, return true
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}

}
