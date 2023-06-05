package validators

import (
	"regexp"
)

// ValidateEmail checks if the given email address is valid
func ValidateEmail(email string) bool {
	// Email validation regular expression
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Match the email against the regular expression
	match, _ := regexp.MatchString(emailRegex, email)

	return match
}
