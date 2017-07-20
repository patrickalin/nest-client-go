# Nest Client in Go

[![Build Status](https://travis-ci.org/patrickalin/nest-client-go.svg?branch=master)](https://travis-ci.org/patrickalin/nest-client-go)
![Build size](https://reposs.herokuapp.com/?path=patrickalin/nest-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/patrickalin/nest-client-go)](https://goreportcard.com/report/github.com/patrickalin/nest-client-go)
[![Coverage Status](https://coveralls.io/repos/github/patrickalin/nest-client-go/badge.svg)](https://coveralls.io/github/patrickalin/nest-client-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Join the chat at https://gitter.im/tockins/nest-client-go](https://badges.gitter.im/tockins/nest-client-go.svg)](https://gitter.im/nest-client-go/)
[![https://img.shields.io/badge/nest-api-go.svg](https://img.shields.io/badge/nest-api-go.svg)](https://github.com/patrickalin/nest-api-go)

A simple Go client for the Nest API.

* It's possible to show informations in the console or in a embedded web server.
* It's also possible to export datas to Time Series Database InfluxData.

## 1] Getting Started

### Prerequisites

* Nest API key (get it here: [Nest api](https://dashboard.nest.com/))

### Installation

* Download the [binary](https://github.com/patrickalin/nest-client-go/releases) for your OS and the [config.yaml](https://github.com/patrickalin/nest-client-go/blob/master/config.yaml) in the same folder.

### Configuration

* Don't forget !!!! : You have to change the API Key in the config.yaml.
* Or use flag "-token xxxxxx"

### Traduction

* This application supports en-us and fr
* Cette application supporte l'anglais et le français

### Binary donwload with config.yaml

| Platform| Architecture | URL|
| ----------| -------- | ------|
|Apple macOS|64-bit Intel| ./goNest-darwin-amd64.bin |
|Linux|64-bit Intel| ./goNest-linux-amd64.bin |
|Windows|64-bit Intel| goNest-windows-amd64.exe |

### Usage with config.yaml or with flag

Execute the binary with the config file "config.yaml" in the same folder.

* Ex : goNest-windows-amd64.exe -token xxxxxxx

There are some others flags : --help for doc

      Usage of ./nest-client-go:
     -debug string
        panic,fatal,error,warning,info,debug
     -devel string
        true,false
     -mock string
        true,false
     -token string
        yourtoken

### Test using Nest Browser

Nest Clientcomes with an embedded web based object browser. Point your web browser by default to `http://localhost:1111/` ensure your server has started successfully.

![Web server](https://raw.githubusercontent.com/patrickalin/nest-client-go/master/img/webserver.png)

### Example : result in the standard console

    ---------------------
    Nest 2017-07-20 17:06:21
    ---------------------
    Ambien TemperatureC :		27.5

### Example : result in a influxData

![InfluxData Image ](https://raw.githubusercontent.com/patrickalin/nest-client-go/master/img/InfluxDB.png)

## Docker Container

docker pull patrickalin/docker-nest
docker run -d  --name=nest -e nestAccessToken=ToBECompleted patrickalin/docker-nest

`https://hub.docker.com/r/patrickalin/docker-nest/`

## 2] Modification code / Compilation

### Pre installation

* install git
* install go from `http://golang.org/`
* If you want install influxData

### Installation env development

    git clone https://github.com/patrickalin/GonestThermostatAPIRest.git
    cd GonestThermostatAPIRest
    export GOPATH=$PWD
    go get -v .
    go build

### Mock

In the config file, you can activate a mock. If you don't have a API key.

* mock: true

### Dev

In the config file, you can change the dev mode to use template, lang locally.

* dev: true

When the dev = false you use assembly files.
Execute "go generate" to refresh assembly files.

### Debug

In the config file, you can change the log level (panic,fatal,error,warn,info,debug)

* logLevel: "debug"

## 3] Thanks

<https://github.com/tixu> for testing and review

<http://mervine.net/json2struct> "transform JSON to Go struct library"

<http://github.com/spf13/viper> "read config library"

## 4] License

The code is licensed under the permissive Apache v2.0 licence. This means you can do what you like with the software, as long as you include the required notices. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.
