package calculator

import (
	"math"

	"github.com/montanaflynn/stats"
)

func StatisticsCalculator(numbers ...int) (int, int, float64) {
	// returns average, mean, median, standard deviation of multiple integer numbers as parameters, as many as you like

	// add the params to a slice
	data := []float64{}

	// if no params, return 0 for all
	if len(numbers) == 0 {
		return 0, 0, 0
	}

	for _, number := range numbers {
		data = append(data, float64(number))
	}

	// // calculate the average
	// var average float64 = stats.medi

	// calculate the mean
	meanFloat, _ := stats.Mean(data)

	// calculate the median
	medianFloat, _ := stats.Median(data)

	// calculate the standard deviation
	standardDeviationFloat, _ := stats.StandardDeviation(data)

	// round the standard deviation to 6 decimal places
	standardDeviationFloat = math.Round(standardDeviationFloat*1000000) / 1000000

	// convert the float64 to int
	var mean int = int(meanFloat)
	var median int = int(medianFloat)

	return mean, median, standardDeviationFloat

}
