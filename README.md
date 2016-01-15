# GoNestThermostatAPIRest
This code execute a call to the Nest API Thermostat in Go (Golang).

#Put the result in the standard console.

     Body :
     {"devices":{"thermostats":{"oJHB1ha6NGOT9493h-fcJY--":{"humidity":35,"locale":"fr-FR","temperature_scale":"C","is_using_emergency_heat":false,"has_fan":false,"software_version":"4.3.3","has_leaf":false,"device_id":"oJHB1ha6NGOT9493h-fcJY--","name":"Dining Room","can_heat":true,"can_cool":false,"hvac_mode":"heat","target_temperature_c":20.0,"target_temperature_f":68,"target_temperature_high_c":24.0,"target_temperature_high_f":75,"target_temperature_low_c":20.0,"target_temperature_low_f":68,"ambient_temperature_c":20.0,"ambient_temperature_f":68,"away_temperature_high_c":24.0,"away_temperature_high_f":76,"away_temperature_low_c":15.0,"away_temperature_low_f":59,"structure_id":"Nhae1XUqlNalBQ82Pfqf","fan_timer_active":false,"name_long":"Dining Room Thermostat","is_online":true,"last_connection":"2015-03-03T08:56:35.754Z"}}},"structures":{"Nhae1XUqlNalBQ82Pfqf":{"name":"Home","country_code":"BE","away":"home","thermostats":["oJHB1ha6NGOT9493h-fcJY--"],"structure_id":"Nhae1XUqlNalBQ82Pfqf"}},"metadata":{"access_token":"c.","client_version":1}}

     Wednesday, 11-Mar-15 09:15:32 CET :> Nest Thermostat Go Call

    DeviceId : 	 	oJHB1ha6NGOT9493h-fcJY-
    SoftwareVersion : 	4.3.3
    Humidity : 	 	45.0
    AmbientTemperatureC : 	18.5
    TargetTemperatureC : 	17.0
    Away : 	 	 	auto-away


     Wednesday, 11-Mar-15 09:15:33 CET :> Send Data to InfluxDB

    DeviceId : 	 	oJHB1ha6NGOT9493h-fcJY-
    SoftwareVersion : 	4.3.3
    Humidity : 	 	45.0
    AmbientTemperatureC : 	18.5
    TargetTemperatureC : 	17.0
    Away : 	 	 	auto-away


     Wednesday, 11-Mar-15 09:16:33 CET :> Send Data to InfluxDB


#Put the result in a influx DB.

![InfluxDB Image ](https://github.com/patrickalin/GoNestThermostatAPIRest/blob/master/img/InfluxDB.png)

After, you can display the result with Grafana

![Grafana Image ](https://github.com/patrickalin/GoNestThermostatAPIRest/blob/master/img/Grafana.png)

#Put data from OpenWeather in InfluxDB

![OpenWeather Image ](https://github.com/patrickalin/GoNestThermostatAPIRest/blob/master/img/OpenWeather.png)

#Pre installation

install git

install go from http://golang.org/

for influxdb, version > 0.9.6

#Installation

    git clone https://github.com/patrickalin/GoNestThermostatAPIRest.git
    cd GoNestThermostatAPIRest
    export GOPATH=$PWD
    go get -v .
    go build

#Configuration

1 You must copy the config.json.example to config.json

    cp config.json.example config.json

2 In the config file modify the secret key receive on https://developer.nest.com/

To test, execute one time :

    curl -L -X GET -H "Accept: application/json" "https://developer-api.nest.com/?auth=c.557"
    with you key

4 Modify all paramameters in config.json

- For InfluxDB isntall the software, the program create database "nest"

#Execution

    ./GoNestThermostatAPIRest

#Debug

In the config file, you can change the log level.

#Thanks

https://github.com/tixu for testing and review

http://mervine.net/json2struct for transform JSON to Go struct

http://github.com/spf13/viper for read config
