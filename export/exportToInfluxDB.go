package export

import (
	"fmt"
	"time"

	config "github.com/patrickalin/GoNestThermostatAPIRest/config"
	nestStructure "github.com/patrickalin/GoNestThermostatAPIRest/nestStructure"

	clientinfluxdb "github.com/influxdata/influxdb/client/v2"

	mylog "github.com/patrickalin/GoMyLog"
)

type influxDBStruct struct {
	Columns []string        `json:"columns"`
	Serie   string          `json:"name"`
	Points  [][]interface{} `json:"points"`
}

type influxDBError struct {
	message error
	advice  string
}

var clientInflux clientinfluxdb.Client

func (e *influxDBError) Error() string {
	return fmt.Sprintf("\n \t InfluxDBError :> %s \n\t InfluxDB Advice:> %s", e.message, e.advice)
}

func sendNestToInfluxDB(oneNest nestStructure.NestStructure, oneConfig config.ConfigStructure) {

	fmt.Printf("\n%s :> Send Nest Data to InfluxDB\n", time.Now().Format(time.RFC850))

	// Create a point and add to batch
	tags := map[string]string{"nest": "living"}
	fields := map[string]interface{}{
		"TargetTemperature":  oneNest.GetTargetTemperatureC(),
		"AmbientTemperature": oneNest.GetAmbientTemperatureC(),
		"Humidity":           oneNest.GetHumidity(),
		"Version":            oneNest.GetSoftwareVersion(),
		"Status":             oneNest.GetAway(),
		"Running":            oneNest.GetAmbientTemperatureF() < oneNest.GetTargetTemperatureF(),
	}

	// Create a new point batch
	bp, err := clientinfluxdb.NewBatchPoints(clientinfluxdb.BatchPointsConfig{
		Database:  oneConfig.InfluxDBDatabase,
		Precision: "s",
	})

	if err != nil {
		mylog.Error.Fatal(&influxDBError{err, "Error sent Data to Influx DB"})
	}

	pt, err := clientinfluxdb.NewPoint("NestData", tags, fields, time.Now())
	bp.AddPoint(pt)

	// Write the batch
	err = clientInflux.Write(bp)

	if err != nil {
		err2 := createDB(oneConfig)
		if err2 != nil {
			mylog.Error.Fatal(&influxDBError{err, "Error with Post : Check if InfluxData is running or if the database nest exists"})
		}
	}

}

func createDB(oneConfig config.ConfigStructure) error {
	fmt.Println("Create Database Nest in InfluxData")

	query := fmt.Sprint("CREATE DATABASE ", oneConfig.InfluxDBDatabase)
	q := clientinfluxdb.NewQuery(query, "", "")

	fmt.Println("Query: ", query)

	_, err := clientInflux.Query(q)
	if err != nil {
		return &influxDBError{err, "Error with : Create database Nest, check if InfluxDB is running"}
	}
	fmt.Println("Database Nest created in InfluxDB")
	return nil
}

func makeClient(oneConfig config.ConfigStructure) (client clientinfluxdb.Client, err error) {
	client, err = clientinfluxdb.NewHTTPClient(
		clientinfluxdb.HTTPConfig{
			Addr:     fmt.Sprintf("http://%s:%s", oneConfig.InfluxDBServer, oneConfig.InfluxDBServerPort),
			Username: oneConfig.InfluxDBUsername,
			Password: oneConfig.InfluxDBPassword,
		})

	if err != nil || client == nil {
		return nil, &influxDBError{err, "Error with creating InfluxDB Client : , check if InfluxDB is running"}
	}
	return client, nil
}

// InitInfluxDB initiate the client influxDB
// Arguments Nest informations, configuration from config file
// Wait events to send to influxDB
func InitInfluxDB(messagesNest chan nestStructure.NestStructure, oneConfig config.ConfigStructure) {

	clientInflux, _ = makeClient(oneConfig)

	go func() {
		mylog.Trace.Println("Receive messagesNest to export InfluxDB")
		for {
			msg := <-messagesNest
			sendNestToInfluxDB(msg, oneConfig)
		}
	}()

}
