package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func BinToHex(bin string) (string, error) {
	// convert bin to hex

	// we are going to split the bin into 4 bits
	var currentNibble string
	var result string

	// detect if the bin input string is empty
	if len(bin) == 0 {
		return "", errors.New("Invalid binary string, no empty strings allowed")
	}

	// first split bin into 4 bits
	for i := len(bin); i >= 0; i -= 4 {
		// convert 4 bits to hex
		if i-4 < 0 {
			// if the length of the current nibble is less than 4
			// add zeros to the front
			zeroString := '0'
			zerosPefix := strings.Repeat(string(zeroString), 4-i)

			currentNibble = zerosPefix + bin[0:i]
		} else {
			// if the length of the current nibble is 4
			// just get the current nibble
			currentNibble = bin[i-4 : i]
		}

		// convert current nibble to decimal
		currentNibbleInt, err := strconv.ParseInt(currentNibble, 2, 0)

		// display parse errors if any
		if err != nil {
			return "", errors.New("Invalid binary string")
		}

		// validate if the last nibble in the loop is zero to break it
		if i <= 0 && currentNibbleInt == 0 {
			break
		}

		// format decimal to hex
		var currentNibbleIntHex string = strconv.FormatInt(currentNibbleInt, 16)

		// append the hex to the result
		result = currentNibbleIntHex + result
	}

	// convert result to uppercase
	result = strings.ToUpper(result)

	// return result
	return result, nil
}

func BinToDec(currentNibble string) {
	panic("unimplemented")
}
