package nestpack

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// read config from config.json
// with the package viper
func GetConfig() (url string, access_token string) {
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	url = viper.GetString("url")
	fmt.Printf("Your URL from config file : %s \n\n", url)

	access_token = viper.GetString("access_token")
	fmt.Printf("You access_token from config file : %s \n\n", access_token)

	return url, access_token
}
