package main

import (
	"nestpack"
)

/*
Get Nest Thermostat Information
*/

func main() {

	// getConfig from the file config.json
	url, access_token := nestpack.GetConfig()
	url = url + access_token

	// get body from Rest API
	rest := nestpack.NewRest()
	body := rest.GetBody(url)

	// get Nest JSON and parse information in Nest Go Structure
	myNest := nestpack.New(body)

	// display major informations to console
	nestpack.DisplayToConsole(myNest)
}
