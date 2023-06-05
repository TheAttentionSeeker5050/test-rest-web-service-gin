package config

import (
	"fmt"
	"os"
)

// Oauth 2 Config function using github
func Oauth2Config() {
	// Oauth Config variables
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("OAUTH_REDIRECT_URL")

	// for now just print the env variables
	fmt.Println("Client ID: ", clientID)
	fmt.Println("Client Secret: ", clientSecret)
	fmt.Println("Redirect URL: ", redirectURL)

}
