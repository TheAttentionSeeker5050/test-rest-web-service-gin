package test

import (
	"testing"
	"workspace/common/calculator"
)

func TestHexToBin(t *testing.T) {

	var has string = calculator.HexToBin("147F")
	var expected string = "1010001111111"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestHexToBinInvalidHexString(t *testing.T) {

	var has string = calculator.BinToHex("147FHIJK")
	var expected string = "error"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestHexToBinCharParam(t *testing.T) {

	var has string = calculator.BinToHex("F")
	var expected string = "15"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}
