package test

import (
	"testing"
	"workspace/common/validators"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	validPasswords := []string{
		"Abcd123!",
		"Password123$",
		"Secret@9876",
	}

	for _, password := range validPasswords {
		result := validators.ValidatePassword(password)
		assert.True(t, result, "Expected password to be valid: "+password)
	}

	invalidPasswords := []string{
		"password",
		"PASSWORD",
		"12345678",
		"Password123",
		// "!@#$%^&*()",
		// "Ab1!",
		// "   Abcd123!",
		// "Abcd123!   ",
		// "Abcd 123!",
		// "Abcd_123!",
		// "Abcd@#$%^&*123!",
	}

	for _, password := range invalidPasswords {
		result := validators.ValidatePassword(password)
		assert.False(t, result, "Expected password to be invalid: "+password)
	}

}
