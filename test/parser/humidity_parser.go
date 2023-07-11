package parser

import (
	"strconv"
)

func ParseHumidity(hex string) (float64, error) {
	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	humidity := float64(value) * 0.5

	return humidity, nil
}
