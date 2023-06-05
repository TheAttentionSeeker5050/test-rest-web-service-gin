package validators

import (
	"regexp"
)

// add password validation method here, it should return true if the password is valid
func ValidatePasswordWithRegex(password string) bool {
	// Invalid password pattern

	// Matches passwords that do not meet the criteria: at least 8 characters, 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character
	pattern := `^(?!.*[A-Z])(?!.*[a-z])(?!.*\d)(?!.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`

	match, _ := regexp.MatchString(pattern, password)
	return !match
}

func ValidatePassword(password string) bool {
	// Password validation criteria
	minLength := 8
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecialChar := false

	// Check the password length
	if len(password) < minLength {
		return false
	}

	// Check for uppercase letter, lowercase letter, digit, and special character
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUppercase = true
		case char >= 'a' && char <= 'z':
			hasLowercase = true
		case char >= '0' && char <= '9':
			hasDigit = true
		case char >= '!' && char <= '/' || char >= ':' && char <= '@' || char >= '[' && char <= '`' || char >= '{' && char <= '~':
			hasSpecialChar = true
		}
	}

	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}
