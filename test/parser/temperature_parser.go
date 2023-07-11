package parser

import (
	"strconv"
)

func ParseTemperature(hex string) (float64, error) {
	f := hex[2:]
	s := hex[:2]
	a := f + s
	value, err := strconv.ParseInt(a, 16, 64)
	
	if err != nil {
		return 0, err
	}

	temperature := float64(value) * 0.1

	return temperature, nil
}
