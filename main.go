package main

import (
	"config"
	"export"
	"fmt"
	"nestStructure"
	"strconv"
	"time"
)

/*
Get Nest Thermostat Information
*/

//name of the config file
const configName = "config"

var myConfig config.ConfigStructure

var nestMessageToConsole = make(chan nestStructure.NestStructure)
var nestMessageToInfluxDB = make(chan nestStructure.NestStructure)

var myTime time.Duration

func main() {

	fmt.Printf("\n %s :> Nest Thermostat Go Call\n\n", time.Now().Format(time.RFC850))

	// getConfig from the file config.json
	myConfig = config.New(configName)
	i, _ := strconv.Atoi(myConfig.RefreshTimer)
	myTime = time.Duration(i) * time.Second

	//init listeners
	if myConfig.ConsoleActivated == "true" {
		export.InitConsole(nestMessageToConsole)
	}
	if myConfig.InfluxDBActivated == "true" {
		export.InitInfluxDB(nestMessageToInfluxDB, myConfig)
	}

	schedule()
}

func schedule() {
	ticker := time.NewTicker(myTime)
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

	// display major informations to console or to influx DB
	nestMessageToConsole <- myNest
	nestMessageToInfluxDB <- myNest

}
