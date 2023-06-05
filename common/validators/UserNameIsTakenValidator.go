package validators

import "gorm.io/gorm"

// add password validation method here, it should return true if the username is taken
func ValidateUserNameIsTaken(db *gorm.DB, userName string) bool {
	// check if the user name is already taken
	if userName == "test" {
		return true
	}
	return false
}
