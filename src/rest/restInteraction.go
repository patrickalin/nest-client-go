package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RestHTTP struct {
	status string
	header http.Header
	body   []byte
}

var debug = false

// Get Rest on the Nest API
func (r *RestHTTP) Get(url string) {

	if debug {
		fmt.Println("Rest Get URL:>", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get URL : " + url)
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if debug {
		fmt.Println("Get response Status:>", resp.Status)
		fmt.Println("Get response Headers:>", resp.Header)
	}

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error with read Body")
		log.Fatal(err)
	}
	if debug {
		fmt.Printf("Body : \n %s \n\n", body)
	}

	if body == nil {
		fmt.Println("Error the body is null, error in the secret key in the config.json ? ")
		log.Fatal(err)
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body
}

// Post Rest on the Nest API
func (r *RestHTTP) PostJSON(url string, buffer []byte) {

	if debug {
		fmt.Println("\n")
		fmt.Println("URL Post :>", url)
		fmt.Printf("Decode Post :> %s \n\n", buffer)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		fmt.Println("Post URL : " + url)
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error with read Body")
		log.Fatal(err)
	}

	if body == nil {
		fmt.Println("Error the body is null, error in the secret key in the config.json ? ")
		log.Fatal(err)
	}

	if resp.Status != "200 OK" || debug {
		fmt.Println("\n URL Post :>", url)
		fmt.Printf("Decode Post :> %s \n\n", buffer)
		fmt.Println("Post response Status:>", resp.Status)
		fmt.Println("Post response Headers:>", resp.Header)
		fmt.Println("Post response Body:>", string(body))
	}

	if resp.Status != "200 OK" {
		fmt.Println("Error status Post Rest ")
		log.Fatal(err)
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body
}

func (r *RestHTTP) GetBody() []byte {
	return r.body
}
