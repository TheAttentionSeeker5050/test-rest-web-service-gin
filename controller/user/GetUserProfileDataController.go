package user

import (
	"net/http"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// the structure for the response body
type ResponseBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

// the api endpoint for retrieving the user profile
func GetUserProfileDataController(c *gin.Context, db *gorm.DB) {
	// prepare the response body
	var responseBody ResponseBody

	// declare the model instance
	var userModel model.UserModel

	// please add the get user profile id data logic from token here

	// get the user model from the context
	modelInstaceResult := model.GetUserModelInstanceById(db, &userModel, 1)

	// check if error
	if modelInstaceResult.Error != nil {
		// return the error
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops! Something went wrong!",
			"error":   modelInstaceResult.Error,
		})
		return
	}

	// create the response body
	responseBody = ResponseBody{
		Username: userModel.UserName,
		Email:    userModel.Email,
		Message:  "User profile retrieved successfully!",
	}

	// return the response
	c.JSON(http.StatusOK, responseBody)
}
