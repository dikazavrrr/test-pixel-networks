package decoder

import (
	"fmt"
	"test/models"
	"test/parser"
	"test/parser/case05"
)

func Decode(hex string) (*models.Result, error) {
	result := &models.Result{}
	F_bat, F_temp, F_hum, F_cnt, F_mag := false, false, false, false, false

	for i := 0; i < len(hex); i += 2 {
		channel := hex[i : i+2]

		if channel == "01" && (i + 6 <= len(hex)) {
			value, err := parser.ParseBattery(hex[i + 4:i + 6])
			if err != nil {
				return nil, err
			}

			result.Battery = value
			F_bat = true
			i += 4
		} else if channel == "03" && (i + 8 <= len(hex)) {
			value, err := parser.ParseTemperature(hex[i + 4:i + 8])
			if err != nil {
				return nil, err
			}

			result.Temperature = value
			F_temp = true
			i += 6
		} else if channel == "04" && (i + 4 <= len(hex)) {
			value, err := parser.ParseHumidity(hex[i + 4:i + 6])
			if err != nil {
				return nil, err
			}

			result.Humidity = value
			F_hum = true
			i += 2
		} else if channel == "05" && (i + 4 <= len(hex)) {
			typeOfChan := hex[i + 2:i + 4]
			var value string
			var err error

			switch typeOfChan {
			case "00":
				value, err = case05.ParseWaterLeakageStatus(hex[i + 4:i + 6])
				if err != nil {
					return nil, err
				}

				i += 4
			case "c8":
				value, err = case05.ParseCounter(hex[i + 4:i + 12])
				if err != nil {
					return nil, err
				}

				i += 10
			default:
				return nil, fmt.Errorf("unknown subtype: %s", typeOfChan)
			}

			result.WaterLeakageStatus = value
			F_cnt = true
		} else if channel == "06" && (i + 4 <= len(hex)) {
			value, err := parser.ParseMagneticStatus(hex[i + 4:i + 6])
			if err != nil {
				return nil, err
			}

			result.MagneticStatus = value
			F_mag = true
			i += 2
		}
	}

	
	return checkFlags(result, F_bat, F_temp, F_hum, F_cnt, F_mag), nil
}
