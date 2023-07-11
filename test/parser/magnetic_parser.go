package parser

import (
	"fmt"
	"strconv"
)

func ParseMagneticStatus(hex string) (string, error) {
	value, err := strconv.ParseInt(hex, 16, 8)
	if err != nil {
		return "", err
	}

	if value == 0 {
		return "Closed", nil
	} else if value == 1 {
		return "Open", nil
	} else {
		return "", fmt.Errorf("unknown magnetic status: %d", value)
	}
}
