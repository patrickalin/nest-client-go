package export

import (
	"fmt"

	"github.com/patrickalin/GoNestThermostatAPIRest/nestStructure"

	mylog "github.com/patrickalin/GoMyLog"
)

//print major informations from a Nest JSON to console
func displayToConsole(oneNest nestStructure.NestStructure) {

	fmt.Printf("\nDeviceId : \t \t%s\n", oneNest.GetDeviceID())
	fmt.Printf("SoftwareVersion : \t%s\n", oneNest.GetSoftwareVersion())
	fmt.Printf("Humidity : \t \t%.1f\n", oneNest.GetHumidity())
	fmt.Printf("AmbientTemperatureC : \t%.1f\n", oneNest.GetAmbientTemperatureC())
	fmt.Printf("TargetTemperatureC : \t%.1f\n", oneNest.GetTargetTemperatureC())
	fmt.Printf("Away : \t \t \t%s\n\n", oneNest.GetAway())
}

//InitConsole listen on the chanel
func InitConsole(messages chan nestStructure.NestStructure) {
	go func() {
		mylog.Trace.Println("Receive message to export Console")

		for {
			msg := <-messages
			displayToConsole(msg)
		}
	}()
}
