package export

import (
	"fmt"
	"nestStructure"
)

var debug = false

//print major informations from a Nest JSON to console
func displayToConsole(oneNest nestStructure.NestStructure) {

	fmt.Printf("\nDeviceId : \t \t%s\n", oneNest.GetDeviceId())
	fmt.Printf("SoftwareVersion : \t%s\n", oneNest.GetSoftwareVersion())
	fmt.Printf("Humidity : \t \t%.1f\n", oneNest.GetHumidity())
	fmt.Printf("AmbientTemperatureC : \t%.1f\n", oneNest.GetAmbientTemperature())
	fmt.Printf("TargetTemperatureC : \t%.1f\n", oneNest.GetTargetTemperature())
	fmt.Printf("Away : \t \t \t%s\n\n", oneNest.GetAway())
}

func InitConsole(messages chan nestStructure.NestStructure) {
	go func() {
		if debug {
			fmt.Println("receive message  to export Console")
		}
		for {
			msg := <-messages
			displayToConsole(msg)
		}
	}()
}
