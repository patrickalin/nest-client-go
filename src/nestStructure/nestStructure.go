package nestStructure

import (
	"config"
	"encoding/json"
	"fmt"
	"log"
	"rest"
)

var debug = false

// generate by http://mervine.net/json2struct
// you must replace your ThermostatID and you structure ID
type nestStructure struct {
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
		StructureID struct {
			Away        string   `json:"away"`
			CountryCode string   `json:"country_code"`
			Name        string   `json:"name"`
			StructureID string   `json:"structure_id"`
			Thermostats []string `json:"thermostats"`
		} `json:"Nhae1XUqlNalBQ82Pfqf6NEt8rObgjPJgNJyoSL6iahQ92AblzZVZw"`
	} `json:"structures"`
}

type NestStructure interface {
	GetDeviceId() string
	GetSoftwareVersion() string
	GetAmbientTemperature() float64
	GetTargetTemperature() float64
	GetHumidity() float64
	GetAway() string
	ShowPrettyAll() int
}

type nestError struct {
	message error
	advice  string
}

func (e *nestError) Error() string {
	return fmt.Sprintf("\n \t NestError :> %s \n\t Advice :> %s", e.message, e.advice)
}

func (nestInfo nestStructure) ShowPrettyAll() int {
	out, err := json.Marshal(nestInfo)
	if err != nil {
		fmt.Println("Error with parsing Json")
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("Decode:> \n %s \n\n", out)

	}
	return 2
}

func (nestInfo nestStructure) GetDeviceId() string {
	return nestInfo.Devices.Thermostats.ThermostatID.DeviceID
}

func (nestInfo nestStructure) GetSoftwareVersion() string {
	return nestInfo.Devices.Thermostats.ThermostatID.SoftwareVersion
}

func (nestInfo nestStructure) GetAmbientTemperature() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.AmbientTemperatureC
}

func (nestInfo nestStructure) GetTargetTemperature() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.TargetTemperatureC
}

func (nestInfo nestStructure) GetHumidity() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.Humidity
}

func (nestInfo nestStructure) GetAway() string {
	return nestInfo.Structures.StructureID.Away
}

func MakeNew(oneConfig config.ConfigStructure) NestStructure {

	// get body from Rest API
	myRest := rest.MakeNew()
	err := myRest.Get(oneConfig.NestURL)

	if err != nil {
		log.Fatal(&nestError{err, "Problem with call rest"})
	}

	var nestInfo nestStructure
	body, _ := myRest.GetBody()
	json.Unmarshal(body, &nestInfo)

	nestInfo.ShowPrettyAll()

	return nestInfo
}
