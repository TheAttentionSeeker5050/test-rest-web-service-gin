package controller

import (
	"net/http"
	"workspace/common/calculator"

	"github.com/gin-gonic/gin"
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

func BasicCalcController(c *gin.Context) {
	var requestBody BasicCalcRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err2 := calculator.BasicCalculator(requestBody.Num1, requestBody.Num2, requestBody.Operator)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	responseData := CalculateResponse{Result: result, CalcParams: requestBody}

	c.JSON(http.StatusOK, responseData)

}
