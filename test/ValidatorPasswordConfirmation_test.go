package test

import (
	"testing"
	"workspace/common/validators"

	"github.com/stretchr/testify/assert"
)

// test case for the password confirmation validator matching
func TestValidatePasswordConfirmation(t *testing.T) {

	// first test case password and password confirmation match
	password := "Abcd123!"
	passwordConfirmation := "Abcd123!"

	result := validators.ValidatePasswordMatchesConfirmation(password, passwordConfirmation)

	assert.True(t, result, "Expected password and password confirmation to match")

	// second test case password and password confirmation do not match
	password = "Abcd123!"
	passwordConfirmation = "Abcd1234!"

	result = validators.ValidatePasswordMatchesConfirmation(password, passwordConfirmation)

	assert.False(t, result, "Expected password and password confirmation to not match")

	// third test case password and password confirmation are empty
	password = ""
	passwordConfirmation = ""

	result = validators.ValidatePasswordMatchesConfirmation(password, passwordConfirmation)

	assert.False(t, result, "Expected password and password confirmation to not match")

}
