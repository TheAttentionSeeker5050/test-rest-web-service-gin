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

func TestBinToHexController(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Register the route handler
	router.POST("/api/v1/calculator/bin-to-hex", controller.BinToHexController)

	// Define the request payload
	requestBody1 := controller.BinToHexRequest{
		Bin: "1010",
	}

	requestBody2 := controller.BinToHexRequest{
		Bin: "1010001111111",
	}

	requestBody3 := controller.BinToHexRequest{
		Bin: "1010001111111a",
	}

	jsonBody1, _ := json.Marshal(requestBody1)
	jsonBody2, _ := json.Marshal(requestBody2)
	jsonBody3, _ := json.Marshal(requestBody3)

	// Create a new HTTP POST request with the JSON payload
	req1, _ := http.NewRequest("POST", "/api/v1/calculator/bin-to-hex", bytes.NewBuffer(jsonBody1))
	req1.Header.Set("Content-Type", "application/json")

	req2, _ := http.NewRequest("POST", "/api/v1/calculator/bin-to-hex", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")

	req3, _ := http.NewRequest("POST", "/api/v1/calculator/bin-to-hex", bytes.NewBuffer(jsonBody3))
	req3.Header.Set("Content-Type", "application/json")

	// Perform the request1
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req1)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	var responseData controller.BinToHexResponse
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, "A", responseData.HexResult)

	// Perform the request2
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req2)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, "147F", responseData.HexResult)

	// delete previous response data, because this next request will return in an error
	responseData = controller.BinToHexResponse{}

	// Perform the request3
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req3)

	// Assert the response status code
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Parse the response body
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data is empty and the error is not empty
	assert.Equal(t, "", responseData.HexResult)
	assert.NotEmpty(t, responseData.Error)

}
