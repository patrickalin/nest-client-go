package nestStructure

import (
	"encoding/json"
	"fmt"
	"time"

	config "github.com/patrickalin/GoNestThermostatAPIRest/config"

	mylog "github.com/patrickalin/GoMyLog"
	rest "github.com/patrickalin/GoRest"
)

// generate by http://mervine.net/json2struct
// you must replace your ThermostatID and you structure ID
type nestStructure struct {
	Devices struct {
		Thermostats struct {
			ThermostatID ThermostatID `json:"noNeeded"`
		} `json:"thermostats"`
	} `json:"devices"`
	Metadata   Metadata `json:"metadata"`
	Structures struct {
		StructureID StructureID `json:"noNeeded"`
	} `json:"structures"`
}

type nestStructureShort struct {
	Devices struct {
		Thermostats interface{} `json:"thermostats"`
	} `json:"devices"`
	Metadata   Metadata    `json:"metadata"`
	Structures interface{} `json:"structures"`
}

type StructureID struct {
	Away        string   `json:"away"`
	CountryCode string   `json:"country_code"`
	Name        string   `json:"name"`
	StructureID string   `json:"structure_id"`
	Thermostats []string `json:"thermostats"`
}

type Metadata struct {
	AccessToken   string  `json:"access_token"`
	ClientVersion float64 `json:"client_version"`
}

type ThermostatID struct {
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
}

// NestStructure is the Interface NestStructure
type NestStructure interface {
	GetDeviceID() string
	GetSoftwareVersion() string
	GetAmbientTemperatureC() float64
	GetTargetTemperatureC() float64
	GetAmbientTemperatureF() float64
	GetTargetTemperatureF() float64
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
		mylog.Error.Fatal(err)
	}
	mylog.Trace.Printf("Decode:> \n %s \n\n", out)
	return 0
}

func (nestInfo nestStructureShort) ShowPrettyAll() int {
	out, err := json.Marshal(nestInfo)
	if err != nil {
		fmt.Println("Error with parsing Json")
		mylog.Error.Fatal(err)
	}
	mylog.Trace.Printf("Decode:> \n %s \n\n", out)
	return 0
}

func (nestInfo nestStructure) GetDeviceID() string {
	return nestInfo.Devices.Thermostats.ThermostatID.DeviceID
}

func (nestInfo nestStructure) GetSoftwareVersion() string {
	return nestInfo.Devices.Thermostats.ThermostatID.SoftwareVersion
}

func (nestInfo nestStructure) GetAmbientTemperatureC() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.AmbientTemperatureC
}

func (nestInfo nestStructure) GetTargetTemperatureF() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.TargetTemperatureF
}

func (nestInfo nestStructure) GetAmbientTemperatureF() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.AmbientTemperatureF
}

func (nestInfo nestStructure) GetTargetTemperatureC() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.TargetTemperatureC
}

func (nestInfo nestStructure) GetHumidity() float64 {
	return nestInfo.Devices.Thermostats.ThermostatID.Humidity
}

func (nestInfo nestStructure) GetAway() string {
	return nestInfo.Structures.StructureID.Away
}

// MakeNew calls Nest and get structureNest
func MakeNew(oneConfig config.ConfigStructure) NestStructure {

	var retry = 0
	var err error
	var duration = time.Minute * 5

	// get body from Rest API
	myRest := rest.MakeNew()
	for retry < 5 {
		err = myRest.Get(oneConfig.NestURL)
		if err != nil {
			mylog.Error.Println(&nestError{err, "Problem with call rest, check the URL and the secret ID in the config file"})
			retry++
			time.Sleep(duration)
		}
	}

	if err != nil {
		mylog.Error.Fatal(&nestError{err, "Problem with call rest, check the URL and the secret ID in the config file"})
	}

	var nestInfo nestStructure
	var nestInfoShort nestStructureShort

	body := myRest.GetBody()

	//err = json.Unmarshal(body, &nestInfo)
	err = json.Unmarshal(body, &nestInfoShort)

	if err != nil {
		mylog.Error.Fatal(&nestError{err, "Problem with json to struct, problem in the struct ?"})
	}

	// not prety but works, one uid is use in the structure nest but I don't like that
	// so I found a work around
	listThermostatsInterface := nestInfoShort.Devices.Thermostats
	listThermostatsMaps := listThermostatsInterface.(map[string]interface{})

	var oneThermostat ThermostatID

	for _, value := range listThermostatsMaps {
		jsonString, _ := json.Marshal(value)
		json.Unmarshal(jsonString, &oneThermostat)
		nestInfo.Devices.Thermostats.ThermostatID = oneThermostat
	}

	listStructuresInterface := nestInfoShort.Structures
	listStructuresMaps := listStructuresInterface.(map[string]interface{})

	var oneStructure StructureID

	for _, value := range listStructuresMaps {
		jsonString, _ := json.Marshal(value)
		json.Unmarshal(jsonString, &oneStructure)
		nestInfo.Structures.StructureID = oneStructure
	}

	nestInfo.Metadata = nestInfoShort.Metadata
	// not prety but works end

	nestInfo.ShowPrettyAll()

	return nestInfo
}
