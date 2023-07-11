package parser

import (
	"strconv"
)

func ParseBattery(hex string) (float64, error) {
	value, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	battery := float64(value)
	return battery, nil
}
