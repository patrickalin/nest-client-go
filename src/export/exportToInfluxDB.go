package export

import (
	"config"
	"encoding/json"
	"fmt"
	"nestStructure"
	"rest"
	"time"
)

type influxDBStruct struct {
	Columns [3]string     `json:"columns"`
	Serie   string        `json:"name"`
	Points  [1][3]float64 `json:"points"`
}

func SendToInfluxDB(oneNest nestStructure.Nest, oneConfig config.ConfigStructure) {

	fmt.Printf("\n %s :> Send Data to InfluxDB\n", time.Now().Format(time.RFC850))

	influxDBData := influxDBStruct{}
	influxDBData.Columns = [3]string{"TargetTemperature", "AmbientTemperature", "Humidity"}
	influxDBData.Points = [1][3]float64{{0.0, 0.0, 0.0}}
	influxDBData.Serie = "NestData"

	influxDBData.Points[0][0] = oneNest.GetTargetTemperature()
	influxDBData.Points[0][1] = oneNest.GetAmbientTemperature()
	influxDBData.Points[0][2] = oneNest.GetHumidity()

	sendPost(influxDBData, oneConfig)

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

func sendPost(influxDBData influxDBStruct, oneConfig config.ConfigStructure) {
	data, _ := json.Marshal(influxDBData)

	data = append(data, byte(']'))
	data = append([]byte(`[`), data...)

	fullURL := fmt.Sprint("http://", oneConfig.InfluxDBServer, ":", oneConfig.InfluxDBServerPort, "/db/", oneConfig.InfluxDBDatabase, "/series?u=", oneConfig.InfluxDBUsername, "&p=", oneConfig.InfluxDBPassword)

	//curl -X POST -d '[{"name":"foo","columns":["val"],"points":[[23]]}]' 'http://localhost:8086/db/nest/series?u=root&p=root'
	oneRest := new(rest.RestHTTP)
	oneRest.PostJSON(fullURL, data)
}
