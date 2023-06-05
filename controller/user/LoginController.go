package user

import (
	"net/http"
	"workspace/common/validators"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginRequest - the request body for the login request
type LoginRequest struct {
	EmailOrUsername string `json:"user"`
	Password        string `json:"password"`
}

// LoginResponse - the response body for the login request
type LoginResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// this is the api endpoint for the user login
func LoginUserController(c *gin.Context, db *gorm.DB) {

	// get the request body
	var requestBody LoginRequest

	// bind the request body to the request params structure
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			LoginResponse{
				Message: "Oops! Something went wrong!",
				Error:   err.Error(),
			},
		)
		return
	}

	// run the authentication logic

	// first check if the credential is using email or username
	// if true it is email, if false it is username
	isEmail := validators.ValidateEmail(requestBody.EmailOrUsername)

	// usermodel instance
	var user model.UserModel

	// if the credential is using email
	if isEmail {
		// Authenticate using email
		db.Where("email = ?", requestBody.EmailOrUsername).First(&user)
	} else {
		// Authenticate using username
		db.Where("username = ?", requestBody.EmailOrUsername).First(&user)
	}

	if user.Id == 0 {
		// User not found, handle the error
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Message: "Oops! Something went wrong!",
			Error:   "Your user or password is incorrect",
		})
		return
	}

	// Compare the password with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(requestBody.Password)); err != nil {
		// Password does not match, handle the error
		c.JSON(http.StatusUnauthorized, LoginResponse{
			Message: "Oops! Something went wrong!",
			Error:   "Your user or password is incorrect",
		})
		return
	}

	//  authenticate the user using OAuth2 using the github strategy
	// 1. redirect the user to the github login page

	// 2. after the user has logged in, github will redirect the user to the callback url
	// 3. the callback url will then exchange the authorization code for an access token
	// 4. the access token will then be used to make API requests or retrieve user information
	// 5. the access token will then be used to authenticate the user

}
