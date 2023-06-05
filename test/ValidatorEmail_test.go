package test

import (
	"testing"
	"workspace/common/validators"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	// Test valid email addresses
	validEmails := []string{
		"test@example.com",
		"john.doe@example.com",
		"123@example.com",
	}

	for _, email := range validEmails {
		result := validators.ValidateEmail(email)
		assert.True(t, result, "Expected email '%s' to be valid, but it was invalid", email)
	}

	// Test invalid email addresses
	invalidEmails := []string{
		"test",
		"test@example",
		"test@example.",
		"test@.com",
	}

	for _, email := range invalidEmails {
		result := validators.ValidateEmail(email)
		assert.False(t, result, "Expected email '%s' to be invalid, but it was valid", email)
	}
}
