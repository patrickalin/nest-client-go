package main

import (
	"config"
	"export"
	"nestStructure"
)

/*
Get Nest Thermostat Information
*/

//name of the config file
const configName = "config"

func main() {

	// getConfig from the file config.json
	myConfig := config.New(configName)

	// get Nest JSON and parse information in Nest Go Structure
	myNest := nestStructure.New(myConfig)
	myNest.ShowPrettyAll()

	if myConfig.ConsoleActivated == "true" {
		// display major informations to console
		export.DisplayToConsole(myNest)
	}

	if myConfig.InfluxDBActivated == "true" {
		// send to influxDB
		export.SendToInfluxDB(myNest, myConfig)
	}
}