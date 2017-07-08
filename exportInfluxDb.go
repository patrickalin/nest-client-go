package main

import (
	"context"
	"fmt"
	"time"

	clientinfluxdb "github.com/influxdata/influxdb/client/v2"
	nest "github.com/patrickalin/nest-api-go"
)

type client struct {
	in       chan nest.Nest
	c        clientinfluxdb.Client
	database string
}

func (c *client) sendnestToInfluxDB(onenest nest.Nest) {

	fmt.Printf("\n%s :> Send nest Data to InfluxDB\n", time.Now().Format(time.RFC850))

	// Create a point and add to batch
	tags := map[string]string{"nest": onenest.GetCity()}
	fields := map[string]interface{}{
		"NumOfFollowers":        onenest.GetNumOfFollowers(),
		"Humidity":              onenest.GetHumidity(),
		"Uv":                    onenest.GetIndexUV(),
		"PressureHpa":           onenest.GetPressureHPa(),
		"PressureInHg":          onenest.GetPressureInHg(),
		"Night":                 onenest.IsNight(),
		"Rain":                  onenest.IsRain(),
		"RainDailyIn":           onenest.GetRainDailyIn(),
		"RainDailyMm":           onenest.GetRainDailyMm(),
		"RainIn":                onenest.GetRainIn(),
		"RainMm":                onenest.GetRainMm(),
		"RainRateIn":            onenest.GetRainRateIn(),
		"RainRateMm":            onenest.GetRainRateMm(),
		"ustainedWindSpeedkmh":  onenest.GetSustainedWindSpeedkmh(),
		"SustainedWindSpeedMph": onenest.GetSustainedWindSpeedMph(),
		"SustainedWindSpeedMs":  onenest.GetSustainedWindSpeedMs(),
		"WindDirection":         onenest.GetWindDirection(),
		"WindGustkmh":           onenest.GetWindGustkmh(),
		"WindGustMph":           onenest.GetWindGustMph(),
		"WindGustMs":            onenest.GetWindGustMs(),
		"TemperatureCelsius":    onenest.GetTemperatureCelsius(),
		"TemperatureFahrenheit": onenest.GetTemperatureFahrenheit(),
		"TimeStamp":             onenest.GetTimeStamp(),
	}

	// Create a new point batch
	bp, err := clientinfluxdb.NewBatchPoints(clientinfluxdb.BatchPointsConfig{
		Database:  c.database,
		Precision: "s",
	})

	if err != nil {
		log.Errorf("Error sent Data to Influx DB : %v", err)
	}

	pt, err := clientinfluxdb.NewPoint("nestData", tags, fields, time.Now())
	bp.AddPoint(pt)

	// Write the batch
	err = c.c.Write(bp)

	if err != nil {
		err2 := c.createDB(c.database)
		if err2 != nil {
			log.Errorf("Check if InfluxData is running or if the database nest exists : %v", err)
		}
	}
}

func (c *client) createDB(InfluxDBDatabase string) error {
	fmt.Println("Create Database nest in InfluxData")

	query := fmt.Sprint("CREATE DATABASE ", InfluxDBDatabase)
	q := clientinfluxdb.NewQuery(query, "", "")

	fmt.Println("Query: ", query)

	_, err := c.c.Query(q)
	if err != nil {
		return fmt.Errorf("Error with : Create database nest, check if InfluxDB is running : %v", err)
	}
	fmt.Println("Database nest created in InfluxDB")
	return nil
}

func initClient(messagesnest chan nest.Nest, InfluxDBServer, InfluxDBServerPort, InfluxDBUsername, InfluxDBPassword, InfluxDatabase string) (*client, error) {
	c, err := clientinfluxdb.NewHTTPClient(
		clientinfluxdb.HTTPConfig{
			Addr:     fmt.Sprintf("http://%s:%s", InfluxDBServer, InfluxDBServerPort),
			Username: InfluxDBUsername,
			Password: InfluxDBPassword,
		})

	if err != nil || c == nil {
		return nil, fmt.Errorf("Error creating database nest, check if InfluxDB is running : %v", err)
	}
	cl := &client{c: c, in: messagesnest, database: InfluxDatabase}
	//need to check how to verify that the db is running
	err = cl.createDB(InfluxDatabase)
	checkErr(err, funcName(), "impossible to create DB", InfluxDatabase)
	return cl, nil
}

// InitInfluxDB initiate the client influxDB
// Arguments nest informations, configuration from config file
// Wait events to send to influxDB
func (c *client) listen(context context.Context) {

	go func() {
		log.Info("Receive messagesnest to export InfluxDB")
		for {
			msg := <-c.in
			c.sendnestToInfluxDB(msg)
		}
	}()
}
