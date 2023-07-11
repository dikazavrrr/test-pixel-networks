package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"test/decoder"
)

func main() {
	fmt.Println("Input:")
	reader := bufio.NewReader(os.Stdin)
	hexString, _ := reader.ReadString('\n')
	hexString = strings.TrimSpace(hexString)
	
	result, err := decoder.Decode(hexString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if result.Battery != 0 {
		fmt.Printf("Battery: %.1f°C\n", result.Battery)
	}

	if result.Temperature != 0 {
		fmt.Printf("Temperature: %.1f°C\n", result.Temperature)
	}

	if result.Humidity != 0 {
		fmt.Printf("Humidity: %.1f%%\n", result.Humidity)
	}

	if result.WaterLeakageStatus != "" {
		if reflect.TypeOf(result).Kind() == reflect.String {
			fmt.Println("WaterLeakageStatus:", result.WaterLeakageStatus)
		} else {
			fmt.Println(result.WaterLeakageStatus)
		}
	}

	if result.MagneticStatus != "" {
		fmt.Println("MagneticStatus:", result.MagneticStatus)
	}
}
