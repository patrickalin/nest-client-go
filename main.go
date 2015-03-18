package main

import (
	"config"
	"export"
	"flag"
	"fmt"
	"mylog"
	"nestStructure"
	"openweathermap"
	"strconv"
	"time"
)

/*
Get Nest Thermostat Information
*/

//name of the config file
const configName = "config"

var (
	nestMessageToConsole            = make(chan nestStructure.NestStructure)
	nestMessageToInfluxDB           = make(chan nestStructure.NestStructure)
	openWeathermapMessageToInfluxDB = make(chan openweathermap.OpenweatherStruct)

	myTime time.Duration

	myConfig config.ConfigStructure

	debug = flag.String("debug", "", "Error=1, Warning=2, Info=3, Trace=4")
)

func main() {

	flag.Parse()

	fmt.Printf("\n %s :> Nest Thermostat Go Call\n\n", time.Now().Format(time.RFC850))

	mylog.Init(mylog.ERROR)

	// getConfig from the file config.json
	myConfig = config.New(configName)

	if *debug != "" {
		myConfig.LogLevel = *debug
	}

	level, _ := strconv.Atoi(myConfig.LogLevel)
	mylog.Init(mylog.Level(level))

	i, _ := strconv.Atoi(myConfig.RefreshTimer)
	myTime = time.Duration(i) * time.Second

	//init listeners
	if myConfig.ConsoleActivated == "true" {
		export.InitConsole(nestMessageToConsole)
	}
	if myConfig.InfluxDBActivated == "true" {
		export.InitInfluxDB(nestMessageToInfluxDB, openWeathermapMessageToInfluxDB, myConfig)
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

	go func() {
		// display major informations to console or to influx DB

		if myConfig.ConsoleActivated == "true" {
			nestMessageToConsole <- myNest
		}
	}()

	go func() {
		// display major informations to console to influx DB
		if myConfig.InfluxDBActivated == "true" {
			nestMessageToInfluxDB <- myNest
		}
	}()

	go func() {
		if myConfig.OpenWeatherActivated == "true" {
			myOpenWeathermap, err := openweathermap.MakeNew(myConfig)
			if err == nil {
				openWeathermapMessageToInfluxDB <- myOpenWeathermap
			}
		}
	}()

}
