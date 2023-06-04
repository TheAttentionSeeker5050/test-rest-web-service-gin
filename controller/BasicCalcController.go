package controller

import (
	"fmt"
	"net/http"
	"workspace/common/calculator"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// the request params structure
type BasicCalcRequest struct {
	Num1     float64 `json:"num1"`
	Num2     float64 `json:"num2"`
	Operator string  `json:"operator"`
}

type CalculateResponse struct {
	Result     float64          `json:"result"`
	CalcParams BasicCalcRequest `json:"calculatorInputs"`
}

func BasicCalcController(c *gin.Context, db *gorm.DB) {
	var requestBody BasicCalcRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err2 := calculator.BasicCalculator(requestBody.Num1, requestBody.Num2, requestBody.Operator)

	if err2 != nil {
		c.JSON(401, gin.H{"error": err2.Error()})
		return
	}

	// save the result to the database

	// create the response data
	var calcHistoryModel model.CalculatorHistoryModel

	// first convert the result to string

	// add the data to the model and convert the result abd request body struct contents to string
	calcHistoryModel.UserName = "anonymous"
	calcHistoryModel.CalculatorType = "BasicCalculator"
	calcHistoryModel.Params = fmt.Sprintf("%.2f", requestBody.Num1) + ", " + fmt.Sprintf("%.2f", requestBody.Num2) + ", " + requestBody.Operator
	// calcHistoryModel.Results = fmt.Sprintf("%f", other.RoundFloat(result, 2))
	calcHistoryModel.Results = fmt.Sprintf("%.2f", result)

	// save the model to the database
	model.CreateCalculatorHistoryModelInstance(db, &calcHistoryModel)

	responseData := CalculateResponse{Result: result, CalcParams: requestBody}

	c.JSON(http.StatusOK, responseData)

}
