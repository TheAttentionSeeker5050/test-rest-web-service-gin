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

func TestHexToBinController(t *testing.T) {

	// Create a new Gin router
	router := gin.Default()

	// Register the route handler
	router.POST("/api/v1/calculator/hex-to-bin", controller.HexToBinController)

	// Define the request payload
	requestBody1 := controller.HexToBinRequest{
		Hex: "A",
	}

	requestBody2 := controller.HexToBinRequest{
		Hex: "AF",
	}

	// this request body should return an error
	requestBody3 := controller.HexToBinRequest{
		Hex: "AFG",
	}

	jsonBody1, _ := json.Marshal(requestBody1)
	jsonBody2, _ := json.Marshal(requestBody2)
	jsonBody3, _ := json.Marshal(requestBody3)

	// Create a new HTTP POST request with the JSON payload
	req1, _ := http.NewRequest("POST", "/api/v1/calculator/hex-to-bin", bytes.NewBuffer(jsonBody1))
	req1.Header.Set("Content-Type", "application/json")

	req2, _ := http.NewRequest("POST", "/api/v1/calculator/hex-to-bin", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")

	req3, _ := http.NewRequest("POST", "/api/v1/calculator/hex-to-bin", bytes.NewBuffer(jsonBody3))
	req3.Header.Set("Content-Type", "application/json")

	// Perform the request1
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req1)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	var responseData controller.HexToBinResponse
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, "1010", responseData.BinResult)

	// Perform the request2
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req2)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, "10101111", responseData.BinResult)

	// delete previous response data, because this next request will return in an error
	responseData = controller.HexToBinResponse{}

	// Perform the request3
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req3)

	// Assert the response status code
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Parse the response body
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data - the result should be empty and the error should not be empty
	assert.Equal(t, "", responseData.BinResult)
	assert.NotEmpty(t, responseData.Error)

}
