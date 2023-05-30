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
