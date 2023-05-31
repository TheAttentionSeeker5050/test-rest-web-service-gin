package controller

import (
	"net/http"
	"workspace/common/calculator"

	"github.com/gin-gonic/gin"
)

// add the request params structure
type HexToBinRequest struct {
	Hex string `json:"hex"`
}

// add the response params structure
type HexToBinResponse struct {
	BinResult  string          `json:"result"`
	CalcParams HexToBinRequest `json:"converterInputs"`
}

// the api handler controller function
func HexToBinController(c *gin.Context) {
	var requestBody HexToBinRequest

	// then bind the request body to the struct and check for errors
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// then call the bin to hex function from the calculator package and get the result
	result, err2 := calculator.HexToBin(requestBody.Hex)

	// if there were errors with the calculator common function, return the error response
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	// then create the response data struct
	responseData := HexToBinResponse{BinResult: result, CalcParams: requestBody}

	// return the response data using json format api response
	c.JSON(http.StatusOK, responseData)

}
