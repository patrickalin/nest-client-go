package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Decode Nest JSON
*/

func main() {

	// push your access_token
	const access_token = "c."
	const url = "https://developer-api.nest.com/?auth="

	// generate by http://mervine.net/json2struct
	// you must replace your ThermostatID
	type NestStructure struct {
		Devices struct {
			Thermostats struct {
				ThermostatID struct {
					AmbientTemperatureC    float64 `json:"ambient_temperature_c"`
					AmbientTemperatureF    float64 `json:"ambient_temperature_f"`
					AwayTemperatureHighC   float64 `json:"away_temperature_high_c"`
					AwayTemperatureHighF   float64 `json:"away_temperature_high_f"`
					AwayTemperatureLowC    float64 `json:"away_temperature_low_c"`
					AwayTemperatureLowF    float64 `json:"away_temperature_low_f"`
					CanCool                bool    `json:"can_cool"`
					CanHeat                bool    `json:"can_heat"`
					DeviceID               string  `json:"device_id"`
					FanTimerActive         bool    `json:"fan_timer_active"`
					HasFan                 bool    `json:"has_fan"`
					HasLeaf                bool    `json:"has_leaf"`
					Humidity               float64 `json:"humidity"`
					HvacMode               string  `json:"hvac_mode"`
					IsOnline               bool    `json:"is_online"`
					IsUsingEmergencyHeat   bool    `json:"is_using_emergency_heat"`
					LastConnection         string  `json:"last_connection"`
					Locale                 string  `json:"locale"`
					Name                   string  `json:"name"`
					NameLong               string  `json:"name_long"`
					SoftwareVersion        string  `json:"software_version"`
					StructureID            string  `json:"structure_id"`
					TargetTemperatureC     float64 `json:"target_temperature_c"`
					TargetTemperatureF     float64 `json:"target_temperature_f"`
					TargetTemperatureHighC float64 `json:"target_temperature_high_c"`
					TargetTemperatureHighF float64 `json:"target_temperature_high_f"`
					TargetTemperatureLowC  float64 `json:"target_temperature_low_c"`
					TargetTemperatureLowF  float64 `json:"target_temperature_low_f"`
					TemperatureScale       string  `json:"temperature_scale"`
				} `json:"oJHB1ha6NGOT9493h-fcJY--gS80WzmN"`
			} `json:"thermostats"`
		} `json:"devices"`
		Metadata struct {
			AccessToken   string  `json:"access_token"`
			ClientVersion float64 `json:"client_version"`
		} `json:"metadata"`
		Structures struct {
			StrcutureID struct {
				Away        string   `json:"away"`
				CountryCode string   `json:"country_code"`
				Name        string   `json:"name"`
				StructureID string   `json:"structure_id"`
				Thermostats []string `json:"thermostats"`
			} `json:"Nhae1XUqlNalBQ82Pfqf6NEt8rObgjPJgNJyoSL6iahQ92AblzZVZw"`
		} `json:"structures"`
	}

	resp, err := http.Get(url + access_token)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error2")
		log.Fatal(err)
	}
	fmt.Printf("Body : \n %s \n\n", body)

	NestDecode := &NestStructure{}
	json.Unmarshal(body, &NestDecode)
	out, err := json.Marshal(NestDecode)
	fmt.Printf("Decode : \n %s \n\n", out)

	fmt.Println("DeviceId : " + NestDecode.Devices.Thermostats.ThermostatID.DeviceID)
	fmt.Printf("AmbientTemperatureC : \t %.1f\n", NestDecode.Devices.Thermostats.ThermostatID.AmbientTemperatureC)
	fmt.Printf("Humidity : \t %.1f\n", NestDecode.Devices.Thermostats.ThermostatID.Humidity)
	fmt.Printf("SoftwareVersion : \t %s\n", NestDecode.Devices.Thermostats.ThermostatID.SoftwareVersion)
	fmt.Printf("TargetTemperatureC : \t %.1f\n", NestDecode.Devices.Thermostats.ThermostatID.TargetTemperatureC)
	fmt.Printf("Away : \t %s\n", NestDecode.Structures.StrcutureID.Away)
}
