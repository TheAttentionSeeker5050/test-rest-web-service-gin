package other

import (
	"math"
)

// RoundFloat rounds a float number to the specified precision
func RoundFloat(number float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(number*scale) / scale
}
