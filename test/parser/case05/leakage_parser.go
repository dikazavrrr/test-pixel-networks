package case05

import (
	"fmt"
)

func ParseWaterLeakageStatus(hex string) (string, error) {
	value := hex[len(hex)-2:]
	if value == "00" {
		return "Not water leaked", nil
	} else if value == "01" {
		return "Water is leaked", nil
	} else {
		return "", fmt.Errorf("unknown status: %s", value)
	}
}
