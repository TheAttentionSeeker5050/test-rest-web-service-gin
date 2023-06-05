package user

import (
	"net/http"
	"workspace/common/validators"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// the request params structure
type RegisterRequest struct {
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	PassWord        string `json:"password"`
	PassWordConfirm string `json:"password_confirm"`
}

// the response data structure
type RegisterResponse struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

// this is the api endpoint for the user registration
func RegisterUserController(c *gin.Context, db *gorm.DB) {

	// get the request body
	var requestBody RegisterRequest

	// bind the request body to the request params structure
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    err.Error(),
			},
		)
		return
	}

	// check if the password and password confirm are the same using the validator package
	passwordsMatch := validators.ValidatePasswordMatchesConfirmation(requestBody.PassWord, requestBody.PassWordConfirm)

	// if passwords do not match
	if passwordsMatch == false {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    "Passwords do not match!",
			},
		)
		return
	}

	// validate if the password is strong enough
	passwordIsStrongEnough := validators.ValidatePassword(requestBody.PassWord)

	// if password is not strong enough
	if passwordIsStrongEnough == false {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    "Password is not strong enough!\nPlease use at least 8 characters, 1 uppercase, 1 lowercase, 1 number and 1 special character!",
			},
		)
		return
	}

	// check if the user name is already taken
	userNameIsTaken := validators.ValidateUserNameIsTaken(db, requestBody.UserName)

	// if user name is already taken
	if userNameIsTaken == true {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    "User name is already taken!",
			},
		)
		return
	}

	// check if the email is already taken
	emailIsTaken := validators.ValidateEmailIsTaken(db, requestBody.Email)

	// if email is already taken
	if emailIsTaken == true {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    "Email is already taken!",
			},
		)
		return
	}

	// validate if email is valid
	emailIsValid := validators.ValidateEmail(requestBody.Email)

	// if email is not valid
	if emailIsValid == false {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    "Email is not valid!",
			},
		)
		return
	}

	// if everything is ok, create the user
	// create the user model instance

	var userModelInstance model.UserModel

	// set the user model instance values
	userModelInstance.UserName = requestBody.UserName
	userModelInstance.Email = requestBody.Email
	userModelInstance.PassWord = requestBody.PassWord

	// create the user
	result := model.CreateUserModelInstance(db, &userModelInstance)

	// check if the user was created
	if result.Error != nil {
		c.JSON(http.StatusBadRequest,
			// a custom response data structure for request error
			RegisterResponse{
				UserName: "",
				Message:  "Oops! Something went wrong!",
				Error:    result.Error.Error(),
			},
		)
		return
	}

	// if everything is ok, return the user data
	c.JSON(http.StatusOK,
		// a custom response data structure for request error
		RegisterResponse{
			UserName: userModelInstance.UserName,
			Message:  "User created successfully!",
			Error:    "",
		},
	)
	return

}
