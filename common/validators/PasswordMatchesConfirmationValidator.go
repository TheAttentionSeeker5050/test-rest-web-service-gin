package validators

// this validator confirms that the password matches the password confirmation
func ValidatePasswordMatchesConfirmation(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation && password != ""
	// && passwordConfirmation != ""
}
