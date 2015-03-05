package main

import (
	"config"
	"rest"
	"export"
	"nestStructure"
)

/*
Get Nest Thermostat Information
*/

func main() {

	// getConfig from the file config.json
	url, access_token := config.GetConfig()
	url = url + access_token

	// get body from Rest API
	restV := rest.NewRest()
	body := restV.GetBody(url)

	// get Nest JSON and parse information in Nest Go Structure
	myNest := nestStructure.New(body)
	myNest.ShowPrettyAll()

	// display major informations to console
	export.DisplayToConsole(myNest)
}
