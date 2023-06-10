package user

import (
	"net/http"
	"time"
	"workspace/common/other"
	"workspace/common/validators"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	Token   string `json:"token"`
}

type LoginResponseError struct {
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
			LoginResponseError{
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
		c.JSON(http.StatusUnauthorized, LoginResponseError{
			Message: "Oops! Something went wrong!",
			Error:   "Your user or password is incorrect",
		})
		return
	}

	// Compare the password with the hashed password stored in the database
	if model.CompareAndValidateHashedPassword(user.PassWord, requestBody.Password) == false {
		// Password does not match, handle the error
		c.JSON(http.StatusUnauthorized, LoginResponseError{
			Message: "Oops! Something went wrong!",
			Error:   "Your user or password is incorrect",
		})
		return
	}

	// generate an access token based on user id
	// please add the access token generation logic here
	tokenString, token := other.GenerateToken(user)

	// check if return is blank, this means that there is an error
	if tokenString == "" || token == nil {
		// return the error
		c.JSON(http.StatusInternalServerError,
			LoginResponseError{
				Message: "Oops! Something went wrong with your session validation",
				Error:   "Authentication Error. Please try logging in again",
			})
		return
	}

	// we will set a cookie for the token
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * time.Duration(token.Claims.(jwt.MapClaims)["exp"].(int64))),
	})

	// return the response
	c.JSON(http.StatusOK, LoginResponse{
		Message: "Login successful!",
		Token:   tokenString,
	})

}
