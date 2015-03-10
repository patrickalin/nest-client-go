# GoNestThermostatAPIRest
How to execute a Get on the Nest API Thermostat in Go (Golang)

#Configuration

You must modify the bin/config.json 
with the secret key receive on https://developer.nest.com/

You must modify the src/nestStructure/nestStructure.go
with your thermostatID end with your structureID

To find these 2 ID's, execute one time :

    curl -L -X GET -H "Accept: application/json" "https://developer-api.nest.com/?auth=c.557" with you key

#Pre installation

install git 

install go from http://golang.org/ 

#Installation

    cd [PATH]
    go get -v .
    export GOBIN=[PATH]/bin
    go install

#Execution

    ./bin/GoNestThermostatAPIRest

# Example of result

     Body : 
     {"devices":{"thermostats":{"oJHB1ha6NGOT9493h-fcJY--":{"humidity":35,"locale":"fr-FR","temperature_scale":"C","is_using_emergency_heat":false,"has_fan":false,"software_version":"4.3.3","has_leaf":false,"device_id":"oJHB1ha6NGOT9493h-fcJY--","name":"Dining Room","can_heat":true,"can_cool":false,"hvac_mode":"heat","target_temperature_c":20.0,"target_temperature_f":68,"target_temperature_high_c":24.0,"target_temperature_high_f":75,"target_temperature_low_c":20.0,"target_temperature_low_f":68,"ambient_temperature_c":20.0,"ambient_temperature_f":68,"away_temperature_high_c":24.0,"away_temperature_high_f":76,"away_temperature_low_c":15.0,"away_temperature_low_f":59,"structure_id":"Nhae1XUqlNalBQ82Pfqf","fan_timer_active":false,"name_long":"Dining Room Thermostat","is_online":true,"last_connection":"2015-03-03T08:56:35.754Z"}}},"structures":{"Nhae1XUqlNalBQ82Pfqf":{"name":"Home","country_code":"BE","away":"home","thermostats":["oJHB1ha6NGOT9493h-fcJY--"],"structure_id":"Nhae1XUqlNalBQ82Pfqf"}},"metadata":{"access_token":"c.","client_version":1}} 
     
     Decode : 
     {"devices":{"thermostats":{"oJHB1ha6NGOT9493h-fcJY--":{"ambient_temperature_c":20,"ambient_temperature_f":68,"away_temperature_high_c":24,"away_temperature_high_f":76,"away_temperature_low_c":15,"away_temperature_low_f":59,"can_cool":false,"can_heat":true,"device_id":"oJHB1ha6NGOT9493h-fcJY--","fan_timer_active":false,"has_fan":false,"has_leaf":false,"humidity":35,"hvac_mode":"heat","is_online":true,"is_using_emergency_heat":false,"last_connection":"2015-03-03T08:56:35.754Z","locale":"fr-FR","name":"Dining Room","name_long":"Dining Room Thermostat","software_version":"4.3.3","structure_id":"Nhae1XUqlNalBQ82Pfqf","target_temperature_c":20,"target_temperature_f":68,"target_temperature_high_c":24,"target_temperature_high_f":75,"target_temperature_low_c":20,"target_temperature_low_f":68,"temperature_scale":"C"}}},"metadata":{"access_token":"c.","client_version":1},"structures":{"Nhae1XUqlNalBQ82Pfqf":{"away":"home","country_code":"BE","name":"Home","structure_id":"Nhae1XUqlNalBQ82Pfqf","thermostats":["oJHB1ha6NGOT9493h-fcJY--gS80WzmN"]}}} 
     
    DeviceId : oJHB1ha6NGOT9493h-fcJY--
    AmbientTemperatureC : 	 20.0
    Humidity : 	 35.0
    SoftwareVersion : 	 4.3.3
    TargetTemperatureC : 	 20.0
    Away : 	 home

#Thanks

http://mervine.net/json2struct for transform JSON to Go struct

github.com/spf13/viper for read config 

https://github.com/rossdylan/influxdbc for client influxdb
