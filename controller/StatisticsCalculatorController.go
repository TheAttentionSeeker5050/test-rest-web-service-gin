// a controller for the statistics calculator post request api endpoint

package controller

import (
	"fmt"
	"net/http"
	"workspace/common/calculator"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// first add the request params structure, we add variable params to the request *args or ...int
type StatisticsCalculatorRequest struct {
	Numbers []int `json:"numbers"`
}

// then add the response params structure
type StatisticsCalculatorResponse struct {
	MeanResult   float64                     `json:"mean"`
	MedianResult float64                     `json:"median"`
	StdDevResult float64                     `json:"standardDeviation"`
	CalcParams   StatisticsCalculatorRequest `json:"calculatorInputs"`
	Error        string                      `json:"error"`
}

// the api handler controller function
func StatisticsCalculatorController(c *gin.Context, db *gorm.DB) {
	// first get the request body struct
	var requestBody StatisticsCalculatorRequest

	// then bind the request body to the struct and check for errors
	if err := c.ShouldBindJSON(&requestBody); err != nil {

		responseData := StatisticsCalculatorResponse{
			MeanResult:   0,
			MedianResult: 0,
			StdDevResult: 0,
			CalcParams:   requestBody,
			Error:        err.Error(),
		}

		c.JSON(400, responseData)
		return
	}

	// then call the statistics calculator function from the calculator package and get the result
	mean, median, standardDeviation, err2 := calculator.StatisticsCalculator(requestBody.Numbers...)

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
	calcHistoryModel.CalculatorType = "StatisticsCalculator"
	// print all the numbers in the array to the params using fmt.Sprintf method separated by commas
	calcHistoryModel.Params = fmt.Sprintf("%v", requestBody.Numbers)
	// format the mean, median and standard deviation results to 2 decimal places separated by commas
	calcHistoryModel.Results = fmt.Sprintf("%.2f, %.2f, %.2f", mean, median, standardDeviation)

	// save the model to the database and return error

	model.CreateCalculatorHistoryModelInstance(db, &calcHistoryModel)

	// then create the response data struct
	responseData := StatisticsCalculatorResponse{
		MeanResult:   mean,
		MedianResult: median,
		StdDevResult: standardDeviation,
		CalcParams:   requestBody,
	}

	// return the response data using json format api response
	c.JSON(http.StatusOK, responseData)

}
