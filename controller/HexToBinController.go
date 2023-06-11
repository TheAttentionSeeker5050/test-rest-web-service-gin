package controller

import (
	"net/http"
	"workspace/common/calculator"
	"workspace/common/other"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// add the request params structure
type HexToBinRequest struct {
	Hex string `json:"hex"`
}

// add the response params structure
type HexToBinResponse struct {
	BinResult  string          `json:"result"`
	CalcParams HexToBinRequest `json:"converterInputs"`
	Error      string          `json:"error"`
}

// the api handler controller function
func HexToBinController(c *gin.Context, db *gorm.DB) {
	var requestBody HexToBinRequest

	// then bind the request body to the struct and check for errors
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		// then create the response data struct
		responseData := HexToBinResponse{BinResult: "", CalcParams: requestBody, Error: err.Error()}

		c.JSON(http.StatusBadRequest, responseData)
		return
	}

	// then call the bin to hex function from the calculator package and get the result
	result, err2 := calculator.HexToBin(requestBody.Hex)

	// if there were errors with the calculator common function, return the error response
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	// save the result to the database

	// create the response data
	var calcHistoryModel model.CalculatorHistoryModel

	// first convert the result to string

	// add the data to the model and convert the result abd request body struct contents to string
	calcHistoryModel.UserName = "anonymous"
	calcHistoryModel.CalculatorType = "HexToBinConverter"
	calcHistoryModel.Params = requestBody.Hex
	// calcHistoryModel.Results = fmt.Sprintf("%f", other.RoundFloat(result, 2))
	calcHistoryModel.Results = result

	// use the validate from cookies auth method
	userIsValidated, userId := other.ValidateUserAuthenticatedFromCookies(c)

	// check if the user is validated
	if userIsValidated {
		// get the user model instance
		var userModel model.UserModel

		// get the user model instance by id
		model.GetUserModelInstanceById(db, &userModel, userId)

		// set the username
		calcHistoryModel.UserName = userModel.UserName
	}

	// save the model to the database and return error
	model.CreateCalculatorHistoryModelInstance(db, &calcHistoryModel)

	// then create the response data struct
	responseData := HexToBinResponse{BinResult: result, CalcParams: requestBody}

	// return the response data using json format api response
	c.JSON(http.StatusOK, responseData)

}
