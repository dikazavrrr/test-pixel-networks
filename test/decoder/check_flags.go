package decoder

import (
	"test/models"
)

func checkFlags(result *models.Result, F_bat, F_temp, F_hum, F_cnt, F_mag bool) *models.Result {
	if !F_bat {
		result.Battery = 0
	}

	if !F_temp {
		result.Temperature = 0
	}

	if !F_hum {
		result.Humidity = 0
	}

	if !F_cnt {
		result.WaterLeakageStatus = ""
	}
	
	if !F_mag {
		result.MagneticStatus = ""
	}

	return result
}