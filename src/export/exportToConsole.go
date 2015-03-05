package export

import (
	"fmt"
	"nestStructure"
)

//print major informations from a Nest JSON to console
func DisplayToConsole(oneNest nestStructure.Nest) {

	fmt.Printf("DeviceId : \t \t%s\n", oneNest.GetDeviceId())
	fmt.Printf("SoftwareVersion : \t%s\n", oneNest.GetSoftwareVersion())
	fmt.Printf("Humidity : \t \t%.1f\n", oneNest.GetHumidity())
	fmt.Printf("AmbientTemperatureC : \t%.1f\n", oneNest.GetAmbientTemperature())
	fmt.Printf("TargetTemperatureC : \t%.1f\n", oneNest.GetTargetTemperature())
	fmt.Printf("Away : \t \t%s\n", oneNest.GetAway())

}
