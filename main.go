// Nest application to export Data nest to console or to influxdb.
package main

//go:generate echo Go Generate!
//go:generate ./command/bindata.sh
//go:generate ./command/bindata-assetfs.sh

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	_ "net/http/pprof"

	"github.com/nicksnyder/go-i18n/i18n"
	nest "github.com/patrickalin/nest-api-go"
	"github.com/patrickalin/nest-client-go/assembly"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//configName name of the config file and log file
const (
	configNameFile = "config"
	logFile        = "nest.log"
)

// Configuration is the structure of the config YAML file
//use http://mervine.net/json2struct
type configuration struct {
	consoleActivated   bool
	hTTPActivated      bool
	historyActivated   bool
	hTTPPort           string
	hTTPSPort          string
	influxDBActivated  bool
	influxDBDatabase   string
	influxDBPassword   string
	influxDBServer     string
	influxDBServerPort string
	influxDBUsername   string
	logLevel           string
	nestAccessToken    string
	nestURL            string
	refreshTimer       time.Duration
	mock               bool
	language           string
	translateFunc      i18n.TranslateFunc
	dev                bool
}

var (
	//Version of the code, fill in in compile.sh -ldflags "-X main.Version=`cat VERSION`"
	Version = "No Version Provided"
	//logger
	log = logrus.New()
)

func init() {
	log.Formatter = new(logrus.JSONFormatter)

	err := os.Remove(logFile)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Error("Failed to log to file, using default stderr")
		return
	}
	log.Out = file
}

func main() {

	//Create context
	logDebug(funcName(), "Create context", "")
	myContext, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)
	go func() {
		select {
		case i := <-signalCh:
			logDebug(funcName(), "Receive interrupt", i.String())
			cancel()
			return
		}
	}()

	logrus.WithFields(logrus.Fields{
		"time":    time.Now().Format(time.RFC850),
		"version": Version,
		"config":  configNameFile,
		"fct":     funcName(),
	}).Info("Nest API")

	//Read configuration from config file
	config, err := readConfig(configNameFile)
	if err != nil {
		logWarn(funcName(), "Config file not loaded error we use flag and default value", os.Args[0])
		config.language = "en-us"
		config.influxDBActivated = false
		config.hTTPActivated = true
		config.hTTPPort = ":1111"
		config.hTTPSPort = ":1112"
		config.consoleActivated = true
		config.refreshTimer = time.Duration(60) * time.Second
		config.nestURL = "https://api.nest.com/api/skydata/"
		config.logLevel = "debug"
		config.mock = true
		config.dev = false
	}

	//Read flags
	logDebug(funcName(), "Get flag from command line", "")
	levelF := flag.String("debug", "", "panic,fatal,error,warning,info,debug")
	tokenF := flag.String("token", "", "yourtoken")
	develF := flag.String("devel", "", "true,false")
	mockF := flag.String("mock", "", "true,false")
	flag.Parse()

	if *levelF != "" {
		config.logLevel = *levelF
	}
	if *tokenF != "" {
		config.nestAccessToken = *tokenF
	}
	if *develF != "" {
		config.dev, err = strconv.ParseBool(*develF)
		checkErr(err, funcName(), "error convert string to bol", "")
	}
	if *mockF != "" {
		config.mock, err = strconv.ParseBool(*mockF)
		checkErr(err, funcName(), "error convert string to bol", "")
	}

	// Set Level log
	level, err := logrus.ParseLevel(config.logLevel)
	checkErr(err, funcName(), "Error parse level", "")
	log.Level = level
	logInfo(funcName(), "Level log", config.logLevel)

	// Context
	ctxsch := context.Context(myContext)

	channels := make(map[string]chan nest.Nest)

	// Traduction
	i18n.ParseTranslationFileBytes("lang/en-us.all.json", readFile("lang/en-us.all.json", config.dev))
	checkErr(err, funcName(), "Error read language file check in config.yaml if dev=false", "")
	i18n.ParseTranslationFileBytes("lang/fr.all.json", readFile("lang/fr.all.json", config.dev))
	checkErr(err, funcName(), "Error read language file check in config.yaml if dev=false", "")
	translateFunc, err := i18n.Tfunc(config.language)
	checkErr(err, funcName(), "Problem with loading translate file", "")

	// Console initialisation
	if config.consoleActivated {
		channels["console"] = make(chan nest.Nest)
		c, err := createConsole(channels["console"], translateFunc, config.dev)
		checkErr(err, funcName(), "Error with initConsol", "")
		c.listen(context.Background())
	}

	// InfluxDB initialisation
	if config.influxDBActivated {
		channels["influxdb"] = make(chan nest.Nest)
		c, err := initClient(channels["influxdb"], config.influxDBServer, config.influxDBServerPort, config.influxDBUsername, config.influxDBPassword, config.influxDBDatabase)
		checkErr(err, funcName(), "Error with initClientInfluxDB", "")
		c.listen(context.Background())
	}

	// WebServer initialisation
	var httpServ *httpServer
	if config.hTTPActivated {
		channels["store"] = make(chan nest.Nest)

		st, err := createStore(channels["store"])
		checkErr(err, funcName(), "Error with history create store", "")
		st.listen(context.Background())

		channels["web"] = make(chan nest.Nest)

		httpServ, err = createWebServer(channels["web"], config.hTTPPort, config.hTTPSPort, translateFunc, config.dev, st)
		checkErr(err, funcName(), "Error with initWebServer", "")
		httpServ.listen(context.Background())
	}

	// get nest JSON and parse information in nest Go Structure
	mynest := nest.New(config.nestURL, config.nestAccessToken, config.mock, log)
	//Call scheduler
	schedule(ctxsch, mynest, channels, config.refreshTimer)

	//If signal to close the program
	<-myContext.Done()
	if httpServ.httpServ != nil {
		logDebug(funcName(), "Shutting down webserver", "")
		err := httpServ.httpServ.Shutdown(myContext)
		checkErr(err, funcName(), "Impossible to shutdown context", "")
	}

	logrus.WithFields(logrus.Fields{
		"fct": "main.main",
	}).Debug("Terminated see nest.log")
}

// The scheduler executes each time "collect"
func schedule(myContext context.Context, mynest nest.Nest, channels map[string]chan nest.Nest, refreshTime time.Duration) {
	ticker := time.NewTicker(refreshTime)
	logDebug(funcName(), "Create scheduler", refreshTime.String())

	collect(mynest, channels)
	for {
		select {
		case <-ticker.C:
			collect(mynest, channels)
		case <-myContext.Done():
			logDebug(funcName(), "Stoping ticker", "")
			ticker.Stop()
			for _, v := range channels {
				close(v)
			}
			return
		}
	}
}

//Principal function which one loops each Time Variable
func collect(mynest nest.Nest, channels map[string]chan nest.Nest) {
	logDebug(funcName(), "Parse informations from API nest", "")

	mynest.Refresh()

	//send message on each channels
	for _, v := range channels {
		v <- mynest
	}
}

// ReadConfig read config from config.json with the package viper
func readConfig(configName string) (configuration, error) {

	var conf configuration
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err, funcName(), "Fielpaths", "")
	dir = dir + "/" + configName

	if err := viper.ReadInConfig(); err != nil {
		logWarn(funcName(), "Error loading the config file", dir)
		return conf, err
	}
	logInfo(funcName(), "The config file loaded", dir)

	//TODO#16 find to simplify this section
	conf.nestURL = viper.GetString("NestURL")
	conf.nestAccessToken = viper.GetString("NestAccessToken")
	conf.influxDBDatabase = viper.GetString("InfluxDBDatabase")
	conf.influxDBPassword = viper.GetString("InfluxDBPassword")
	conf.influxDBServer = viper.GetString("InfluxDBServer")
	conf.influxDBServerPort = viper.GetString("InfluxDBServerPort")
	conf.influxDBUsername = viper.GetString("InfluxDBUsername")
	conf.consoleActivated = viper.GetBool("ConsoleActivated")
	conf.influxDBActivated = viper.GetBool("InfluxDBActivated")
	conf.historyActivated = viper.GetBool("historyActivated")
	conf.refreshTimer = time.Duration(viper.GetInt("RefreshTimer")) * time.Second
	conf.hTTPActivated = viper.GetBool("HTTPActivated")
	conf.hTTPPort = viper.GetString("HTTPPort")
	conf.hTTPSPort = viper.GetString("hTTPSPort")
	conf.logLevel = viper.GetString("LogLevel")
	conf.mock = viper.GetBool("mock")
	conf.language = viper.GetString("language")
	conf.dev = viper.GetBool("dev")

	// Check if one value of the structure is empty
	v := reflect.ValueOf(conf)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i)
		//TODO#16
		//v.Field(i).SetString(viper.GetString(v.Type().Field(i).Name))
		if values[i] == "" {
			return conf, fmt.Errorf("Check if the key " + v.Type().Field(i).Name + " is present in the file " + dir)
		}
	}
	if token := os.Getenv("nestAccessToken"); token != "" {
		conf.nestAccessToken = token
	}
	return conf, nil
}

//Read file and return []byte
func readFile(fileName string, dev bool) []byte {
	if dev {
		fileByte, err := ioutil.ReadFile(fileName)
		checkErr(err, funcName(), "Error reading the file", fileName)
		return fileByte
	}

	fileByte, err := assembly.Asset(fileName)
	checkErr(err, funcName(), "Error reading the file", fileName)
	return fileByte
}
