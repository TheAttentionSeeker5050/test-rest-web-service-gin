package controller

import (
	"fmt"
	"reflect"
	"workspace/common/calculator"

	"github.com/gin-gonic/gin"
)

type BasicCalcRequest struct {
	// the basic calculator request params
	Num1     float64 `json:"num1"`
	Num2     float64 `json:"num2"`
	Operator string  `json:"operator"`
}

func BasicCalcController(c *gin.Context) {

	// the request params structure
	var calcRequest BasicCalcRequest
	fmt.Println("bind json function result:--->", c.Bind(&calcRequest))
	// bind the request params to the structure
	if c.BindJSON(&calcRequest) == nil {
		// print the request params
		fmt.Println("num1: type", reflect.TypeOf(calcRequest.Num1), "value", calcRequest.Num1)
		fmt.Println("num2: type", reflect.TypeOf(calcRequest.Num2), "value", calcRequest.Num2)
		fmt.Println("operator: type", reflect.TypeOf(calcRequest.Operator), "value", calcRequest.Operator)

		// calculate the result
		has, err := calculator.BasicCalculator(calcRequest.Num1, calcRequest.Num2, calcRequest.Operator)
		// call the basic calculator function
		if err != nil {
			fmt.Println("Error: ", err)
			c.JSON(422, gin.H{
				"message":    "Error",
				"error":      err,
				"statusCode": 422,
			})
		} else {
			c.JSON(200, gin.H{
				"message":       "Success",
				"result":        has,
				"requestParams": calcRequest, // this idk if its going to work
				"statusCode":    200,
				"date":          "00/00/0000", // some dummy date for now
			})
		}
	} else {

		c.JSON(400, gin.H{
			"message":    "Error",
			"error":      "Invalid request params",
			"statusCode": 400,
		})
	}
}
