package test

import (
	"testing"
	"workspace/common/calculator"
)

func TestBinToHex(t *testing.T) {

	has, err := calculator.BinToHex("1010001111111")
	var expected string = "147F"

	if err != nil {
		t.Errorf("There was an error with the calculation %v ", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexInvalidBinString(t *testing.T) {

	has, err := calculator.BinToHex("1010001111111a")
	var expected string = ""

	if err != nil {
		t.Log("There was an error with the calculation:", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexOneNibbleParam(t *testing.T) {

	has, err := calculator.BinToHex("1010")
	var expected string = "A"

	if err != nil {
		t.Errorf("There was an error with the calculation: %v ", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexOneNibbleParam2(t *testing.T) {

	has, err := calculator.BinToHex("1011")
	var expected string = "B"

	if err != nil {
		t.Errorf("There was an error with the calculation: %v ", err)
	} else if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}
