package controller

import (
	"fmt"
	"workspace/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// the request params structure, please correct this later
// type dataRequest struct {
// 	Num1 float64 `json:"num1"`
// }

// POST Rest API request handler for calculator history
func CalcHistorySaveToDbController(c *gin.Context, db *gorm.DB) {

	// the request params structure from the model package
	var dataRequest model.CalculatorHistoryModel

	// then bind the request body to the struct and check for errors
	if err := c.ShouldBindJSON(&dataRequest); err != nil {
		// then return the error response
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// now call the create instance to database function from the model package
	result := model.CreateCalculatorHistoryModelInstance(db, &dataRequest)

	// check for errors
	if result.Error != nil {
		// then return the error response
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	// then return the response
	c.JSON(200, gin.H{"message": "success"})
}

// GET Rest API request handler for calculator history to return all the saved entries
func CalcHistoryGetAllController(c *gin.Context, db *gorm.DB) {

	// the request params structure from the model package
	var dataRequest []model.CalculatorHistoryModel

	// now call the get all instances from database function from the model package
	result := model.GetAllCalculatorHistoryModelInstances(db, &dataRequest)

	// check for errors
	if result.Error != nil {
		// then return the error response
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	// print the dataRequest
	fmt.Println(dataRequest)

	// then return the response
	c.JSON(200, dataRequest)
}
