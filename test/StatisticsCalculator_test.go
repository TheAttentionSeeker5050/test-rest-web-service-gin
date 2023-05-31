package test

import (
	"testing"
	"workspace/common/calculator"
)

func TestStatisticsCalculator(t *testing.T) {

	mean, median, standardDeviation, err := calculator.StatisticsCalculator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expectedMean, expectedMedian, expectedStandardDeviation := 5.5, 5.5, 2.872281

	if err != nil {
		t.Errorf("There was an error with the calculation: %v ", err)
		return
	}

	if mean != expectedMean {
		t.Errorf("Expected mean %v but got %v", expectedMean, mean)
	} else {
		t.Log("Success", mean)
	}

	if median != expectedMedian {
		t.Errorf("Expected median %v but got %v", expectedMedian, median)
	} else {
		t.Log("Success", median)
	}

	if standardDeviation != expectedStandardDeviation {
		t.Errorf("Expected standard deviation %v but got %v", expectedStandardDeviation, standardDeviation)
	} else {
		t.Log("Success", standardDeviation)
	}
}

func TestStatisticsCalculatorOneParam(t *testing.T) {

	mean, median, standardDeviation, err := calculator.StatisticsCalculator(1)
	expectedMean, expectedMedian, expectedStandardDeviation := 1.0, 1.0, 0.0

	if err != nil {
		t.Errorf("There was an error with the calculation: %v ", err)
		return
	}

	if mean != expectedMean {
		t.Errorf("Expected mean %v but got %v", expectedMean, mean)
	} else {
		t.Log("Success", mean)
	}

	if median != expectedMedian {
		t.Errorf("Expected median %v but got %v", expectedMedian, median)
	} else {
		t.Log("Success", median)
	}

	if standardDeviation != expectedStandardDeviation {
		t.Errorf("Expected standard deviation %v but got %v", expectedStandardDeviation, standardDeviation)
	} else {
		t.Log("Success", standardDeviation)
	}
}

func TestStatisticsCalculatorAllParamsAreZero(t *testing.T) {

	mean, median, standardDeviation, err := calculator.StatisticsCalculator(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	expectedMean, expectedMedian, expectedStandardDeviation := 0.0, 0.0, 0.0

	if err != nil {
		t.Errorf("There was an error with the calculation: %v ", err)
		return
	}

	if mean != expectedMean {
		t.Errorf("Expected mean %v but got %v", expectedMean, mean)
	} else {
		t.Log("Success", mean)
	}

	if median != expectedMedian {
		t.Errorf("Expected median %v but got %v", expectedMedian, median)
	} else {
		t.Log("Success", median)
	}

	if standardDeviation != expectedStandardDeviation {
		t.Errorf("Expected standard deviation %v but got %v", expectedStandardDeviation, standardDeviation)
	} else {
		t.Log("Success", standardDeviation)
	}

}

func TestStatisticsCalculatorNoParams(t *testing.T) {

	mean, median, standardDeviation, err := calculator.StatisticsCalculator()
	expectedMean, expectedMedian, expectedStandardDeviation := 0.0, 0.0, 0.0

	if err == nil {
		t.Log("This test case launched an error, as it should be: ", err)
		return
	}

	if mean != expectedMean {
		t.Errorf("Expected mean %v but got %v", expectedMean, mean)
	} else {
		t.Log("Success", mean)
	}

	if median != expectedMedian {
		t.Errorf("Expected median %v but got %v", expectedMedian, median)
	} else {
		t.Log("Success", median)
	}

	if standardDeviation != expectedStandardDeviation {
		t.Errorf("Expected standard deviation %v but got %v", expectedStandardDeviation, standardDeviation)
	} else {
		t.Log("Success", standardDeviation)
	}
}
