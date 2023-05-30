package test

import (
	"testing"
	"workspace/common/calculator"
)

func TestBinToHex(t *testing.T) {

	var has string = calculator.BinToHex("1010001111111")
	var expected string = "147F"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexInvalidBinString(t *testing.T) {

	var has string = calculator.BinToHex("1010001111111a")
	var expected string = "error"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexOneNibbleParam(t *testing.T) {

	var has string = calculator.BinToHex("1010")
	var expected string = "A"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}

func TestBinToHexOneNibbleParam2(t *testing.T) {

	var has string = calculator.BinToHex("1011")
	var expected string = "B"

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success", has)
	}
}
