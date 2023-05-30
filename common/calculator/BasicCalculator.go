package calculator

import "errors"

func BasicCalculator(num1 float64, num2 float64, operator string) (float64, error) {
	// perform basic calculations
	var result float64 = 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			// it would be better to be able to return an error here
			return 0, errors.New("Cannot divide by zero")
		}

		result = num1 / num2
	default:
		return 0, errors.New("Invalid operator")
	}

	return result, nil
	// return result, nil
}
