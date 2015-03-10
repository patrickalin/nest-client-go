package config

import (
	"fmt"
	"log"
	"os"
    "path/filepath"

	"github.com/spf13/viper"
)

var debug = false

const nest_url = "nest_url"
const nest_access_token = "nest_access_token"
const influxDB_database = "influxDB_database"
const influxDB_password = "influxDB_password"
const influxDB_server = "influxDB_server"
const influxDB_server_port = "influxDB_server_port"
const influxDB_username = "influxDB_username"
const console_activated = "console_activated"
const influxDB_activated = "influxDB_activated"

type ConfigStructure struct {
	ConsoleActivated   string `json:"console_activated"`
	InfluxDBActivated  string `json:"influxDB_activated"`
	InfluxDBDatabase   string `json:"influxDB_database"`
	InfluxDBPassword   string `json:"influxDB_password"`
	InfluxDBServer     string `json:"influxDB_server"`
	InfluxDBServerPort string `json:"influxDB_server_port"`
	InfluxDBUsername   string `json:"influxDB_username"`
	NestAccessToken    string `json:"nest_access_token"`
	NestURL            string `json:"nest_url"`
}

type Config interface {
	GetURL() string
}

// read config from config.json
// with the package viper

func (configInfo ConfigStructure) ReadConfig(configName string) ConfigStructure {
	viper.SetConfigName(configName)

    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }

    if debug {
		fmt.Printf("The config file loaded is :> %s/%s \n \n", dir, configName)
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("File not found:> %s/%s \n \n", dir, configName)
		log.Fatal(err)
	}

	configInfo.NestURL = viper.GetString(nest_url)
	if configInfo.NestURL == "" {
		log.Fatal("Check if the key :> " + nest_url + " is present in the file " + configName)
	}

	configInfo.NestAccessToken = viper.GetString(nest_access_token)
	if configInfo.NestURL == "" {
		log.Fatal("Check if the key :> " + nest_access_token + " is present in the file " + configName)
	}

	configInfo.NestURL += configInfo.NestAccessToken
	if debug {
		fmt.Printf("Your URL from config file :> %s \n\n", configInfo.NestURL)
	}

	configInfo.InfluxDBDatabase = viper.GetString(influxDB_database)
	if configInfo.InfluxDBDatabase == "" {
		log.Fatal("Check if the key :> " + influxDB_database + " is present in the file " + configName)
	}

	configInfo.InfluxDBPassword = viper.GetString(influxDB_password)
	if configInfo.InfluxDBPassword == "" {
		log.Fatal("Check if the key :> " + influxDB_password + " is present in the file " + configName)
	}

	configInfo.InfluxDBServer = viper.GetString(influxDB_server)
	if configInfo.InfluxDBServer == "" {
		log.Fatal("Check if the key :> " + influxDB_server + " is present in the file " + configName)
	}

	configInfo.InfluxDBServerPort = viper.GetString(influxDB_server_port)
	if configInfo.InfluxDBServerPort == "" {
		log.Fatal("Check if the key :> " + influxDB_server_port + " is present in the file " + configName)
	}

	configInfo.InfluxDBUsername = viper.GetString(influxDB_username)
	if configInfo.InfluxDBUsername == "" {
		log.Fatal("Check if the key :> " + influxDB_username + " is present in the file " + configName)
	}

	configInfo.ConsoleActivated = viper.GetString(console_activated)
	if configInfo.ConsoleActivated == "" {
		log.Fatal("Check if the key :> " + console_activated + " is present in the file " + configName)
	}

	configInfo.InfluxDBActivated = viper.GetString(influxDB_activated)
	if configInfo.InfluxDBActivated == "" {
		log.Fatal("Check if the key :> " + influxDB_activated + " is present in the file " + configName)
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
