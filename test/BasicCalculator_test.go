package test

import (
	"fmt"
	"testing"
	"workspace/common/calculator"
)

func TestBasicCalculator(t *testing.T) {

	has, err := calculator.BasicCalculator(1, 2, "+")
	var expected float64 = 3

	if err != nil {
		fmt.Println(err) // This should not be triggered
	}

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBasicCalculatorInvalidOperator(t *testing.T) {

	has, err := calculator.BasicCalculator(1, 2, "a")
	var expected float64 = 0

	if err != nil {
		fmt.Println(err) // Invalid operator
	}

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBasicCalculatorDivideByZero(t *testing.T) {

	has, err := calculator.BasicCalculator(1, 0, "/")
	var expected float64 = 0

	if err != nil {
		fmt.Println(err) // Zero cannot be used 0
	}

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}
