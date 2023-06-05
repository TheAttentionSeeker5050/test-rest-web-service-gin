package validators

import (
	"regexp"
)

// validator function for username
func ValidateUserName(userName string) bool {
	// username should be no longer than 24 characters, no shorter than 6 characters, no special characters, no spaces, start with lette, can contain numbers, special characters _ and - are allowed

	// create the regex pattern
	regexPattern := "^[a-zA-Z][a-zA-Z0-9_-]{5,24}$"

	// compile the regex pattern
	regex, err := regexp.Compile(regexPattern)

	// check if there is an error
	if err != nil {
		// return false if there is an error
		return false
	}

	// return the result of the regex match
	return regex.MatchString(userName)
}
