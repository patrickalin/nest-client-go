package openweathermap

import (
	"config"
	"encoding/json"
	"fmt"
	"mylog"
	"rest"
)

type openweatherStruct struct {
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Cod   float64 `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   float64 `json:"dt"`
	ID   float64 `json:"id"`
	Main struct {
		Humidity float64 `json:"humidity"`
		Pressure float64 `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Rain struct {
		_h float64 `json:"3h"`
	} `json:"rain"`
	Sys struct {
		Country string  `json:"country"`
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
	} `json:"sys"`
	Weather []struct {
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		ID          float64 `json:"id"`
		Main        string  `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   float64 `json:"deg"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type OpenweatherStruct interface {
	GetCity() string
	GetHumidity() float64
	GetPressure() float64
	GetTemp() float64
	GetWindSpeed() float64
	GetWindDeg() float64
	GetSunrise() float64
	GetSunset() float64
	GetDescription() string
}

type openweatherError struct {
	message error
	advice  string
}

func (e *openweatherError) Error() string {
	return fmt.Sprintf("\n \t OpenweatherError :> %s \n\t Advice :> %s", e.message, e.advice)
}

func (openweatherInfo openweatherStruct) ShowPrettyAll() int {
	out, err := json.Marshal(openweatherInfo)
	if err != nil {
		fmt.Println("Error with parsing Json")
		mylog.Error.Fatal(err)
	}
	mylog.Trace.Printf("Decode openweather:> \n %s \n\n", out)
	return 2
}

func (openweatherInfo openweatherStruct) GetCity() string {
	return openweatherInfo.Name
}

func (openweatherInfo openweatherStruct) GetHumidity() float64 {
	return openweatherInfo.Main.Humidity
}

func (openweatherInfo openweatherStruct) GetPressure() float64 {
	return openweatherInfo.Main.Pressure
}

func (openweatherInfo openweatherStruct) GetTemp() float64 {
	return openweatherInfo.Main.Temp
}

func (openweatherInfo openweatherStruct) GetWindSpeed() float64 {
	return openweatherInfo.Wind.Speed
}

func (openweatherInfo openweatherStruct) GetWindDeg() float64 {
	return openweatherInfo.Wind.Deg
}

func (openweatherInfo openweatherStruct) GetSunrise() float64 {
	return openweatherInfo.Sys.Sunrise
}

func (openweatherInfo openweatherStruct) GetSunset() float64 {
	return openweatherInfo.Sys.Sunset
}

func (openweatherInfo openweatherStruct) GetDescription() string {
	return openweatherInfo.Weather[0].Description
}

func MakeNew(oneConfig config.ConfigStructure) (OpenweatherStruct, error) {

	// get body from Rest API
	myRest := rest.MakeNew()
	err := myRest.Get(oneConfig.OpenweathermapURL + oneConfig.OpenweathermapCityID)
	if err != nil {
		fmt.Println(&openweatherError{err, "Problem with call rest Openweather"})
		return nil, err
	}

	var openweatherInfo openweatherStruct
	body := myRest.GetBody()
	json.Unmarshal(body, &openweatherInfo)

	openweatherInfo.ShowPrettyAll()

	return openweatherInfo, nil
}
