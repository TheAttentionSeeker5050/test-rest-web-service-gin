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

func TestBasicCalcController(t *testing.T) {
	// initialize the database
	db, _ := config.MockDBSetup(t)

	db.AutoMigrate(&model.CalculatorHistoryModel{})

	// drop all previous entries from the table
	db.Where("user_id > ?", 0).Delete(&model.CalculatorHistoryModel{})

	// Create a new Gin router
	router := gin.Default()
	// router = routers.CalcRouter(router, db)

	// // Register the route handler
	router.POST("/api/v1/calculator/basic-calc", func(ctx *gin.Context) {
		controller.BasicCalcController(ctx, db)
	})

	// Define the request payload
	requestBody := controller.BasicCalcRequest{
		Num1:     1.0,
		Num2:     2.0,
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
