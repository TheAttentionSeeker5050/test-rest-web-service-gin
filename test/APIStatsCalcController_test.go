package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace/config"
	"workspace/controller"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStatisticsCalculatorController(t *testing.T) {
	// initialize the database
	db, err := config.MockDBSetup(t)
	assert.Nil(t, err)

	db.AutoMigrate(&model.CalculatorHistoryModel{})

	// drop all previous entries from the table
	db.Where("user_id > ?", 0).Delete(&model.CalculatorHistoryModel{})

	// Create a new Gin router
	router := gin.Default()

	// Register the route handler
	router.POST("/api/v1/calculator/statistics-calc", func(ctx *gin.Context) {
		controller.StatisticsCalculatorController(ctx, db)
	})

	// Define the request payload
	requestBody1 := controller.StatisticsCalculatorRequest{
		Numbers: []int{1, 2, 3, 4, 5},
	}

	requestBody2 := controller.StatisticsCalculatorRequest{
		Numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	// this request body should return an error
	requestBody3 := controller.StatisticsCalculatorRequest{
		Numbers: []int{},
	}

	jsonBody1, _ := json.Marshal(requestBody1)
	jsonBody2, _ := json.Marshal(requestBody2)
	jsonBody3, _ := json.Marshal(requestBody3)

	// Create a new HTTP POST request with the JSON payload
	req1, _ := http.NewRequest("POST", "/api/v1/calculator/statistics-calc", bytes.NewBuffer(jsonBody1))
	req1.Header.Set("Content-Type", "application/json")

	req2, _ := http.NewRequest("POST", "/api/v1/calculator/statistics-calc", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")

	req3, _ := http.NewRequest("POST", "/api/v1/calculator/statistics-calc", bytes.NewBuffer(jsonBody3))
	req3.Header.Set("Content-Type", "application/json")

	// Perform the request1
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req1)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	var responseData controller.StatisticsCalculatorResponse
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, float64(3), responseData.MeanResult)
	assert.Equal(t, float64(3), responseData.MedianResult)
	assert.Equal(t, float64(1.414214), responseData.StdDevResult)
	// assert.Equal(t, float64(1.4142135623730951), responseData.StdDevResult)

	// Perform the request2
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req2)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	json.Unmarshal(resp.Body.Bytes(), &responseData)

	// Assert the response data
	assert.Equal(t, float64(5), responseData.MeanResult)
	assert.Equal(t, float64(5), responseData.MedianResult)
	assert.Equal(t, float64(2.581989), responseData.StdDevResult)

	// Perform the request3
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req3)

	// Assert the response status code
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Parse the response body
	var errorResponseData controller.StatisticsCalculatorResponse
	json.Unmarshal(resp.Body.Bytes(), &errorResponseData)

	// Assert the response data
	assert.Equal(t, float64(0), errorResponseData.MeanResult)
	assert.Equal(t, float64(0), errorResponseData.MedianResult)
	assert.Equal(t, float64(0), errorResponseData.StdDevResult)
	assert.NotEmpty(t, errorResponseData.Error)

}
