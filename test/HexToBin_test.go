package test

import (
	"testing"
	"workspace/common/calculator"
)

func TestHexToBin(t *testing.T) {

	has, err := calculator.HexToBin("147F")
	// var expected string = "0001010001111111"
	var expected string = "0001010001111111"

	if err != nil {
		t.Errorf("There was an error with the calculation %v ", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestHexToBinInvalidHexString(t *testing.T) {

	has, err := calculator.HexToBin("147FHIJK")
	var expected string = "error"

	if err != nil {
		t.Log("This test case launched an error, as it should have:", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestHexToBinOneCharParam(t *testing.T) {

	has, err := calculator.HexToBin("F")
	var expected string = "1111"

	if err != nil {
		t.Log("This test case launched an error, as it should have:", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}
