package config

import (
	"fmt"
	"os"
	"path/filepath"

	mylog "github.com/patrickalin/GoMyLog"
	viper "github.com/spf13/viper"
)

const nestURL = "nest_url"
const nestAccessToken = "nest_access_token"
const influxDBDatabase = "influxDB_database"
const influxDBPassword = "influxDB_password"
const influxDBServer = "influxDB_server"
const influxDBServerPort = "influxDB_server_port"
const influxDBUsername = "influxDB_username"
const consoleActivated = "console_activated"
const influxDBActivated = "influxDB_activated"
const refreshTimer = "refresh_timer"
const logLevel = "log_level"

//ConfigStructure is the structure of the config YAML file
//use http://mervine.net/json2struct
type ConfigStructure struct {
	ConsoleActivated   string `json:"console_activated"`
	InfluxDBActivated  string `json:"influxDB_activated"`
	InfluxDBDatabase   string `json:"influxDB_database"`
	InfluxDBPassword   string `json:"influxDB_password"`
	InfluxDBServer     string `json:"influxDB_server"`
	InfluxDBServerPort string `json:"influxDB_server_port"`
	InfluxDBUsername   string `json:"influxDB_username"`
	LogLevel           string `json:"log_level"`
	NestAccessToken    string `json:"nest_access_token"`
	NestURL            string `json:"nest_url"`
	RefreshTimer       string `json:"refresh_timer"`
}

//Config GetURL return the URL from the config file
type Config interface {
	GetURL() string
}

// ReadConfig read config from config.json
// with the package viper
func (configInfo ConfigStructure) ReadConfig(configName string) ConfigStructure {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		mylog.Error.Fatal(err)
	}

	mylog.Trace.Printf("The config file loaded is :> %s/%s \n \n", dir, configName)

	dir = dir + "/" + configName

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("File not found:> %s/%s \n \n", dir, configName)
		mylog.Error.Fatal(err)
	}

	configInfo.NestURL = viper.GetString(nestURL)
	if configInfo.NestURL == "" {
		mylog.Error.Fatal("Check if the key :> " + nestURL + " is present in the file " + dir)
	}

	configInfo.NestAccessToken = os.Getenv("nest_secretid")
	if configInfo.NestAccessToken == "" {
		configInfo.NestAccessToken = viper.GetString(nestAccessToken)
		if configInfo.NestURL == "" {
			mylog.Error.Fatal("Check if the key :> " + nestAccessToken + " is present in the file " + dir)
		}
	}

	configInfo.NestURL += configInfo.NestAccessToken
	mylog.Trace.Printf("Your URL from config file :> %s \n\n", configInfo.NestURL)

	configInfo.InfluxDBDatabase = viper.GetString(influxDBDatabase)
	if configInfo.InfluxDBDatabase == "" {
		mylog.Error.Fatal("Check if the key " + influxDBDatabase + " is present in the file " + dir)
	}

	configInfo.InfluxDBPassword = viper.GetString(influxDBPassword)
	if configInfo.InfluxDBPassword == "" {
		mylog.Error.Fatal("Check if the key " + influxDBPassword + " is present in the file " + dir)
	}

	configInfo.InfluxDBServer = viper.GetString(influxDBServer)
	if configInfo.InfluxDBServer == "" {
		mylog.Error.Fatal("Check if the key " + influxDBServer + " is present in the file " + dir)
	}

	configInfo.InfluxDBServerPort = viper.GetString(influxDBServerPort)
	if configInfo.InfluxDBServerPort == "" {
		mylog.Error.Fatal("Check if the key " + influxDBServerPort + " is present in the file " + dir)
	}

	configInfo.InfluxDBUsername = viper.GetString(influxDBUsername)
	if configInfo.InfluxDBUsername == "" {
		mylog.Error.Fatal("Check if the key " + influxDBUsername + " is present in the file " + dir)
	}

	configInfo.ConsoleActivated = viper.GetString(consoleActivated)
	if configInfo.ConsoleActivated == "" {
		mylog.Error.Fatal("Check if the key " + consoleActivated + " is present in the file " + dir)
	}

	configInfo.InfluxDBActivated = viper.GetString(influxDBActivated)
	if configInfo.InfluxDBActivated == "" {
		mylog.Error.Fatal("Check if the key " + influxDBActivated + " is present in the file " + dir)
	}

	configInfo.RefreshTimer = viper.GetString(refreshTimer)
	if configInfo.RefreshTimer == "" {
		mylog.Error.Fatal("Check if the key " + refreshTimer + " is present in the file " + dir)
	}

	configInfo.LogLevel = viper.GetString(logLevel)
	if configInfo.LogLevel == "" {
		mylog.Error.Fatal("Check if the key " + logLevel + " is present in the file " + dir)
	}

	return configInfo
}

//New create the configStructure and fill in
func New(configName string) ConfigStructure {
	var configInfo ConfigStructure
	configInfo = configInfo.ReadConfig(configName)
	return configInfo
}

// GetURL return nestURL
func (configInfo ConfigStructure) GetURL() string {
	return configInfo.NestURL
}
