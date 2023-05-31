package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBasicCalcController(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register the route handler
	router.POST("/api/v1/calculator/basic-calc", controller.BasicCalcController)

	// Define the request payload
	requestBody := controller.BasicCalcRequest{
		Num1:     1,
		Num2:     2,
		Operator: "+",
	}

	jsonBody, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request with the JSON payload
	req, _ := http.NewRequest("POST", "/api/v1/calculator/basic-calc", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	var responseData controller.CalculateResponse
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, 3.0, responseData.Result)
	assert.Equal(t, requestBody, responseData.CalcParams)

}
