package nestpack

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type restHTTP struct {
	body []byte
}

type RestHTTP interface {
	GetBody(url string) []byte
}

// Get Rest on the Nest API
func (rest *restHTTP) GetBody(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error with Get URL")
		log.Fatal(err)
	}

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error with read Body")
		log.Fatal(err)
	}
	fmt.Printf("Body : \n %s \n\n", body)

	if body == nil {
		fmt.Println("Error the body is null, error in the secret key in the config.json ? ")
	}
	defer resp.Body.Close()

	return body
}

func NewRest() RestHTTP {
	var rest = new(restHTTP)
	return rest
}
