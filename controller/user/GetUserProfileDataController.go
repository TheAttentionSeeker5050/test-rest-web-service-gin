package user

import (
	"fmt"
	"net/http"
	"workspace/common/other"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// the structure for the response body
type ResponseBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

// the structure for the response body if error
type ResponseBodyError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// the api endpoint for retrieving the user profile
func GetUserProfileDataController(c *gin.Context, db *gorm.DB) {
	// prepare the response body
	var responseBody ResponseBody

	// declare the model instance
	var userModel model.UserModel

	// get token string from secure cookie
	tokenBytes, err := other.SecureCookieObj.GetValue(nil, c.Request)

	// check if error on getting the token string
	if err != nil {
		// return the error
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, ResponseBodyError{
			Message: "Oops! Something went wrong with your cookie validation",
			Error:   "Authentication Error. Please try logging in again",
		})
		return
	}

	// verify the token
	tokenClaims, isTokenValid := other.VerifyAccessToken(string(tokenBytes))
	if !isTokenValid {
		// return the error
		c.JSON(http.StatusUnauthorized, ResponseBodyError{
			Message: "Oops! Something went wrong with your session validation",
			Error:   "Authentication Error. Please try logging in again",
		})
		return
	}

	// get the user id from the token claims
	userId := int(tokenClaims["user_id"].(float64))

	// check if the user id is valid
	if userId <= 0 {
		// return the error
		c.JSON(http.StatusUnauthorized, ResponseBodyError{
			Message: "Oops! Something went wrong!",
			Error:   "Internal Error. Could not find user data",
		})
		return
	}

	// get the user model from the context
	modelInstaceResult := model.GetUserModelInstanceById(db, &userModel, userId)

	// check if error
	if modelInstaceResult.Error != nil {
		// return the error
		c.JSON(http.StatusInternalServerError, ResponseBodyError{
			Message: "Oops! Something went wrong!",
			Error:   "Internal Error. Could not find user data",
		})
		return
	}

	// create the response body for successfull request
	responseBody = ResponseBody{
		Username: userModel.UserName,
		Email:    userModel.Email,
		Message:  "User profile retrieved successfully!",
	}

	// return the response
	c.JSON(http.StatusOK, responseBody)
}
