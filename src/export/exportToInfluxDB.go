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
	Columns [3]string     `json:"columns"`
	Serie   string        `json:"name"`
	Points  [1][3]float64 `json:"points"`
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
	influxDBData.Columns = [3]string{"TargetTemperature", "AmbientTemperature", "Humidity"}
	influxDBData.Points = [1][3]float64{{0.0, 0.0, 0.0}}
	influxDBData.Serie = "NestData"

	influxDBData.Points[0][0] = oneNest.GetTargetTemperature()
	influxDBData.Points[0][1] = oneNest.GetAmbientTemperature()
	influxDBData.Points[0][2] = oneNest.GetHumidity()

	err := sendPost(influxDBData, oneConfig)
	if err != nil {
		log.Fatal(&influxDBError{err, "Error sent Data to Influx DB"})
	}
	/*influxDBData.Columns = [1]string{"Version"}
	influxDBData.Points = [1][1]float64{{0.0}}
	influxDBData.Serie = "SoftwareVersion"

	influxDBData.Points[0][0] = oneNest.GetSoftwareVersion()

	sendPost(influxDBData, oneConfig)

	influxDBData.Columns = [1]string{"Status"}
	influxDBData.Points = [1][1]float64{{0.0}}
	influxDBData.Serie = "Away"

	influxDBData.Points[0][0] = oneNest.GetAway()

	sendPost(influxDBData, oneConfig)*/
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
