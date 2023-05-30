package test

import (
	"testing"
	"workspace/common"
)

func TestExample(t *testing.T) {

	var has int = common.ExampleFunction(1, 1)
	var expected int = 2

	if has != expected {
		t.Errorf("Expected %v but got %v", expected, has)
	} else {
		t.Log("Success")
	}
}
