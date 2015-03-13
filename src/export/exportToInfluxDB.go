package export

import (
	"config"
	"encoding/json"
	"fmt"
	"log"
	"nestStructure"
	"rest"
	"time"
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
	return fmt.Sprintf("\n \t InfluxDBError :> %s \n\t Advice :> %s", e.message, e.advice)
}

func sendToInfluxDB(oneNest nestStructure.NestStructure, oneConfig config.ConfigStructure) {

	fmt.Printf("\n %s :> Send Data to InfluxDB\n", time.Now().Format(time.RFC850))

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
		log.Fatal(&influxDBError{err, "Error sent Data to Influx DB"})
	}

}

func sendPost(influxDBData influxDBStruct, oneConfig config.ConfigStructure) (err error) {
	data, _ := json.Marshal(influxDBData)

	data = append(data, byte(']'))
	data = append([]byte(`[`), data...)

	fullURL := fmt.Sprint("http://", oneConfig.InfluxDBServer, ":", oneConfig.InfluxDBServerPort, "/db/", oneConfig.InfluxDBDatabase, "/series?u=", oneConfig.InfluxDBUsername, "&p=", oneConfig.InfluxDBPassword)

	//curl -X POST -d '[{"name":"foo","columns":["val"],"points":[[23]]}]' 'http://localhost:8086/db/nest/series?u=root&p=root'
	oneRest := rest.MakeNew()
	err = oneRest.PostJSON(fullURL, data)
	if err != nil {
		return &influxDBError{err, "Error with Post : Check if InfluxDB is running"}
	}
	return nil
}

func InitInfluxDB(messages chan nestStructure.NestStructure, oneConfig config.ConfigStructure) {
	go func() {
		if debug {
			fmt.Println("receive message  to export Console")
		}
		for {
			msg := <-messages
			sendToInfluxDB(msg, oneConfig)
		}
	}()
}
