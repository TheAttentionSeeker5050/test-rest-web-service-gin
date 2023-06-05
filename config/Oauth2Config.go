package config

import (
	"fmt"
	"os"
)

// Oauth Config variables
var clientID string = os.Getenv("OAUTH_CLIENT_ID")
var clientSecret string = os.Getenv("OAUTH_CLIENT_SECRET")
var redirectURL string = os.Getenv("OUATH_REDIRECT_URL")

// Oauth 2 Config function using github
func Oauth2Config() {
	// for now just print the env variables
	fmt.Println("Client ID: ", clientID)
	fmt.Println("Client Secret: ", clientSecret)
	fmt.Println("Redirect URL: ", redirectURL)
}
