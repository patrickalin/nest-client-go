package config

import (
	"fmt"
	"os"
	"path/filepath"

	mylog "github.com/patrickalin/GoMyLog"
	viper "github.com/spf13/viper"
)

const nest_url = "nest_url"
const nest_access_token = "nest_access_token"
const influxDB_database = "influxDB_database"
const influxDB_password = "influxDB_password"
const influxDB_server = "influxDB_server"
const influxDB_server_port = "influxDB_server_port"
const influxDB_username = "influxDB_username"
const console_activated = "console_activated"
const influxDB_activated = "influxDB_activated"
const refresh_timer = "refresh_timer"
const log_level = "log_level"

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

type Config interface {
	GetURL() string
}

// read config from config.json
// with the package viper

func (configInfo ConfigStructure) ReadConfig(configName string) ConfigStructure {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		mylog.Error.Fatal(err)
	}

	mylog.Trace.Print("The config file loaded is :> %s/%s \n \n", dir, configName)

	dir = dir + "/" + configName

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("File not found:> %s/%s \n \n", dir, configName)
		mylog.Error.Fatal(err)
	}

	configInfo.NestURL = viper.GetString(nest_url)
	if configInfo.NestURL == "" {
		mylog.Error.Fatal("Check if the key :> " + nest_url + " is present in the file " + dir)
	}

	configInfo.NestAccessToken = viper.GetString(nest_access_token)
	if configInfo.NestURL == "" {
		mylog.Error.Fatal("Check if the key :> " + nest_access_token + " is present in the file " + dir)
	}

	configInfo.NestURL += configInfo.NestAccessToken
	mylog.Trace.Printf("Your URL from config file :> %s \n\n", configInfo.NestURL)

	configInfo.InfluxDBDatabase = viper.GetString(influxDB_database)
	if configInfo.InfluxDBDatabase == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_database + " is present in the file " + dir)
	}

	configInfo.InfluxDBPassword = viper.GetString(influxDB_password)
	if configInfo.InfluxDBPassword == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_password + " is present in the file " + dir)
	}

	configInfo.InfluxDBServer = viper.GetString(influxDB_server)
	if configInfo.InfluxDBServer == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_server + " is present in the file " + dir)
	}

	configInfo.InfluxDBServerPort = viper.GetString(influxDB_server_port)
	if configInfo.InfluxDBServerPort == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_server_port + " is present in the file " + dir)
	}

	configInfo.InfluxDBUsername = viper.GetString(influxDB_username)
	if configInfo.InfluxDBUsername == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_username + " is present in the file " + dir)
	}

	configInfo.ConsoleActivated = viper.GetString(console_activated)
	if configInfo.ConsoleActivated == "" {
		mylog.Error.Fatal("Check if the key " + console_activated + " is present in the file " + dir)
	}

	configInfo.InfluxDBActivated = viper.GetString(influxDB_activated)
	if configInfo.InfluxDBActivated == "" {
		mylog.Error.Fatal("Check if the key " + influxDB_activated + " is present in the file " + dir)
	}

	configInfo.RefreshTimer = viper.GetString(refresh_timer)
	if configInfo.RefreshTimer == "" {
		mylog.Error.Fatal("Check if the key " + refresh_timer + " is present in the file " + dir)
	}

	configInfo.LogLevel = viper.GetString(log_level)
	if configInfo.LogLevel == "" {
		mylog.Error.Fatal("Check if the key " + log_level + " is present in the file " + dir)
	}

	return configInfo
}

func New(configName string) ConfigStructure {
	var configInfo ConfigStructure
	configInfo = configInfo.ReadConfig(configName)
	return configInfo
}

func (configInfo ConfigStructure) GetURL() string {
	return configInfo.NestURL
}