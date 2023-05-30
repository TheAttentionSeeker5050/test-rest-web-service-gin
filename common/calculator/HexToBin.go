package calculator

import (
	"strconv"
	"strings"
)

func HexToBin(hex string) string {
	// convert a hex string to bin

	// name the result variable
	var result string = ""

	// convert hex to decimal
	for i := 0; i < len(hex); i++ {
		// convert current hex to decimal
		currentHexInt, err := strconv.ParseInt(string(hex[i]), 16, 0)

		// display parse errors if any
		if err != nil {
			return "error"
		}

		// convert decimal to bin
		var currentHexIntBin string = strconv.FormatInt(currentHexInt, 2)

		// add zeros to the front if the length of the current bin is less than 4
		if len(currentHexIntBin) < 4 {
			zeroString := '0'
			zerosPrefix := strings.Repeat(string(zeroString), 4-len(currentHexIntBin))
			currentHexIntBin = zerosPrefix + currentHexIntBin
		}

		// append the bin to the result
		result = result + currentHexIntBin
	}

	// dec, err := strconv.ParseInt(bin, 16, 0)
	return result

}
