package case05

import (
	"fmt"
	"strconv"
)

func ParseCounter(hex string) (string, error) {
	value, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Counter: %d", value), nil
}