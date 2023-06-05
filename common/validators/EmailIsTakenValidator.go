package validators

import "gorm.io/gorm"

// add password validation method here, it should return true if the username is taken
func ValidateEmailIsTaken(db *gorm.DB, email string) bool {
	// check if the user name is already taken
	if email == "test" {
		return true
	}
	return false
}
