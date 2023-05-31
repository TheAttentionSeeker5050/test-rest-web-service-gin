// controller for the bin to hex post request api endpoint
package controller

import (
	"net/http"
	"workspace/common/calculator"

	"github.com/gin-gonic/gin"
)

// first add the request params structure
type BinToHexRequest struct {
	Bin string `json:"bin"`
}

// then add the response params structure
type BinToHexResponse struct {
	HexResult  string          `json:"hexResult"`
	CalcParams BinToHexRequest `json:"converterInputs"`
}

// the api handler controller function
func BinToHexController(c *gin.Context) {
	// first get the request body struct
	var requestBody BinToHexRequest

	// then bind the request body to the struct and check for errors
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// then call the bin to hex function from the calculator package and get the result
	result := calculator.BinToHex(requestBody.Bin)

	// errors should go here but for the moment the function doesnt return them

	// then create the response data struct
	responseData := BinToHexResponse{HexResult: result, CalcParams: requestBody}

	// return the response data using json format api response
	c.JSON(http.StatusOK, responseData)

}
