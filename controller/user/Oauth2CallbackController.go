package user

import (
	"net/http"
	"workspace/config"

	"context"

	"github.com/gin-gonic/gin"
)

// OAuthCallbackController handles the callback URL after OAuth 2.0 authorization
func OAuthCallbackController(c *gin.Context) {
	// Get the authorization code from the query parameters
	code := c.Query("code")

	// Exchange the authorization code for an access token
	token, err := config.Oauth2Config().Exchange(context.Background(), code)
	if err != nil {
		// Handle error case
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to exchange authorization code for access token",
		})
		return
	}

	// Use the access token to make API requests or retrieve user information
	// Example: Get the user's email from the API
	client := config.Oauth2Config().Client(context.Background(), token)
	response, err := client.Get("https://api.github.com/user")
	// response, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		// Handle error case
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user information",
		})
		return
	}
	defer response.Body.Close()

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "OAuth authentication successful",
	})
}
