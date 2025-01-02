package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floatSlice := make([]float64, len(strings))

	for stringIndex, stringValue := range strings {
		floatPrice, error := strconv.ParseFloat(stringValue, 64)

		if error != nil {
			fmt.Println("An error occurred while reading slice lines: ", error)
			return nil, errors.New("an error occurred while reading slice lines")
		}
		floatSlice[stringIndex] = floatPrice
	}

	return floatSlice, nil
}
