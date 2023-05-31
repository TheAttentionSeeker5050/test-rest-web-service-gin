package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNoRouteAPI(t *testing.T) {
	// we test routes that are not defined in the router

	// Create a new Gin router
	router := gin.Default()

	// Register the route handler for no route
	router.NoRoute(controller.NoRouteController)

	// Create 3 new HTTP GET requests for registered routes
	req1, _ := http.NewRequest("GET", "/api/v1/calculator/blabla", nil)
	req2, _ := http.NewRequest("GET", "/api/v1/blablabla", nil)
	req3, _ := http.NewRequest("GET", "/blablabla", nil)

	// Perform the request1
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req1)

	// Assert the response status code
	assert.Equal(t, http.StatusNotFound, resp.Code)

	// assert that the error message on the json response is not empty
	assert.NotEmpty(t, resp.Body.String())

	// Perform the request2
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req2)

	// Assert the response status code
	assert.Equal(t, http.StatusNotFound, resp.Code)

	// assert that the error message on the json response is not empty
	assert.NotEmpty(t, resp.Body.String())

	// Perform the request3
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req3)

	// Assert the response status code
	assert.Equal(t, http.StatusNotFound, resp.Code)

	// assert that the error message on the json response is not empty
	assert.NotEmpty(t, resp.Body.String())
}
