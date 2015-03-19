package export

import (
	"config"
	"encoding/json"
	"fmt"
	"nestStructure"
	"time"

	mylog "github.com/patrickalin/GoMyLog"
	rest "github.com/patrickalin/GoRest"
)

type influxDBStruct struct {
	Columns [6]string         `json:"columns"`
	Serie   string            `json:"name"`
	Points  [1][6]interface{} `json:"points"`
}

type influxDBError struct {
	message error
	advice  string
}

func (e *influxDBError) Error() string {
	return fmt.Sprintf("\n \t InfluxDBError :> %s \n\t InfluxDB Advice:> %s %s %s %s", e.message, e.advice)
}

func sendNestToInfluxDB(oneNest nestStructure.NestStructure, oneConfig config.ConfigStructure) {

	fmt.Printf("\n %s :> Send Nest Data to InfluxDB\n", time.Now().Format(time.RFC850))

	influxDBData := influxDBStruct{}
	influxDBData.Columns = [6]string{"TargetTemperature", "AmbientTemperature", "Humidity", "Version", "Status", "Running"}

	pts := [1][6]float64{{0.0, 0.0, 0.0, 0.0, 0.0, 0.0}}

	for i, d := range pts {
		influxDBData.Points[0][i] = interface{}(d)
	}

	influxDBData.Serie = "NestData"

	influxDBData.Points[0][0] = oneNest.GetTargetTemperatureC()
	influxDBData.Points[0][1] = oneNest.GetAmbientTemperatureC()
	influxDBData.Points[0][2] = oneNest.GetHumidity()
	influxDBData.Points[0][3] = oneNest.GetSoftwareVersion()
	influxDBData.Points[0][4] = oneNest.GetAway()
	influxDBData.Points[0][5] = oneNest.GetAmbientTemperatureF() < oneNest.GetTargetTemperatureF()

	err := sendPost(influxDBData, oneConfig)
	if err != nil {
		mylog.Error.Fatal(&influxDBError{err, "Error sent Data to Influx DB"})
	}

}

func sendPost(influxDBData interface{}, oneConfig config.ConfigStructure) (err error) {
	data, _ := json.Marshal(influxDBData)

	data = append(data, byte(']'))
	data = append([]byte(`[`), data...)

	fullURL := fmt.Sprint("http://", oneConfig.InfluxDBServer, ":", oneConfig.InfluxDBServerPort, "/db/", oneConfig.InfluxDBDatabase, "/series?u=", oneConfig.InfluxDBUsername, "&p=", oneConfig.InfluxDBPassword)

	//curl -X POST -d '[{"name":"foo","columns":["val"],"points":[[23]]}]' 'http://localhost:8086/db/nest/series?u=root&p=root'
	oneRest := rest.MakeNew()
	err = oneRest.PostJSON(fullURL, data)
	if err != nil {
		err2 := createDB(oneConfig)
		if err2 != nil {
			return &influxDBError{err2, "Error with Post : Check if InfluxDB is running or if the database nest exists"}
		}
	}
	return nil
}

func createDB(oneConfig config.ConfigStructure) error {
	type createDB struct {
		Name string `json:"name"`
	}

	fmt.Println("\n Create Database Nest\n")

	nestDB := createDB{}
	fullURL := fmt.Sprint("http://", oneConfig.InfluxDBServer, ":", oneConfig.InfluxDBServerPort, "/db?u=", oneConfig.InfluxDBUsername, "&p=", oneConfig.InfluxDBPassword)
	nestDB.Name = oneConfig.InfluxDBDatabase
	data, _ := json.Marshal(nestDB)

	oneRest := rest.MakeNew()
	err := oneRest.PostJSON(fullURL, data)
	if err != nil {
		return &influxDBError{err, "Error with Post : create database Nest"}
	}
	return nil
}

func InitInfluxDB(messagesNest chan nestStructure.NestStructure, oneConfig config.ConfigStructure) {

	go func() {
		mylog.Trace.Println("receive messagesNest  to export InfluxDB")
		for {
			msg := <-messagesNest
			sendNestToInfluxDB(msg, oneConfig)
		}
	}()

}
