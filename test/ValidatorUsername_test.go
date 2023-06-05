package test

import (
	"testing"
	"workspace/common/validators"

	"github.com/stretchr/testify/assert"
)

func TestValidateUsername(t *testing.T) {
	// username should be no longer than 24 characters, no shorter than 6 characters, no special characters, no spaces, start with lette, can contain numbers, special characters _ and - are allowed

	// Test valid usernames
	validUsernames := []string{
		"test555",
		"test-555",
		"test_555",
		// with capital letters
		"Test555",
		"Test-555",
		"TeSS_555",
	}

	for _, username := range validUsernames {
		result := validators.ValidateUserName(username)
		assert.True(t, result, "Expected username '%s' to be valid, but it was invalid", username)
	}

	// Test invalid usernames
	invalidUsernames := []string{
		// too short
		"tes",
		// too long (24 characters+1)
		"test555test555test555test55t5",
		// with spaces
		"test 555",
		// with special characters
		"test@555",
		// not starting with letter
		"555test",
	}

	for _, username := range invalidUsernames {
		result := validators.ValidateUserName(username)
		assert.False(t, result, "Expected username '%s' to be invalid, but it was valid", username)
	}
}
