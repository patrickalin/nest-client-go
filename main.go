package main

import (
	"config"
	"export"
	"nestStructure"
	"time"
	"fmt"
)

/*
Get Nest Thermostat Information
*/

//name of the config file
const configName = "config"
var myConfig config.ConfigStructure

func main() {

	fmt.Printf("\n %s :> Nest Thermostat Go Call\n", time.Now().Format(time.RFC850))

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
        case <- ticker.C:
            repeat()
        case <- quit:
            ticker.Stop()
            return
        }
    }
}


func repeat() {
	// get Nest JSON and parse information in Nest Go Structure
	myNest := nestStructure.New(myConfig)

	if myConfig.ConsoleActivated == "true" {
		// display major informations to console
		export.DisplayToConsole(myNest)
	}

	if myConfig.InfluxDBActivated == "true" {
		// send to influxDB
		export.SendToInfluxDB(myNest, myConfig)
	}
}