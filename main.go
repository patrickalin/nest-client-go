package main

import (
	"config"
	"export"
	"fmt"
	"nestStructure"
	"time"
)

/*
Get Nest Thermostat Information
*/

//name of the config file
const configName = "config"

var myConfig config.ConfigStructure

func main() {

	fmt.Printf("\n %s :> Nest Thermostat Go Call\n\n", time.Now().Format(time.RFC850))

	// getConfig from the file config.json
	myConfig = config.New(configName)

	schedule()
}

func schedule() {
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	repeat()
	for {
		select {
		case <-ticker.C:
			repeat()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func repeat() {
	// get Nest JSON and parse information in Nest Go Structure
	myNest := nestStructure.MakeNew(myConfig)

	if myConfig.ConsoleActivated == "true" {
		// display major informations to console
		export.DisplayToConsole(myNest)
	}

	if myConfig.InfluxDBActivated == "true" {
		// send to influxDB
		export.SendToInfluxDB(myNest, myConfig)
	}
}
