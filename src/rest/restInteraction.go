package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type restHTTP struct {
	status string
	header http.Header
	body   []byte
}

type RestHTTP interface {
	GetBody() ([]byte, error)
	Get(url string) (err error)
	PostJSON(url string, buffer []byte) (err error)
}

type restError struct {
	message error
	url     string
	advice  string
}

func (e *restError) Error() string {
	return fmt.Sprintf("\n \t RestError :> %s \n\t URL :> %s \n\t Advice :> %s", e.message, e.url, e.advice)
}

func MakeNew() (rest RestHTTP) {
	return &restHTTP{}
}

var debug = false

// Get Rest on the Nest API
func (r *restHTTP) Get(url string) (err error) {

	if debug {
		fmt.Println("Rest Get URL:>", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return &restError{err, url, "Check your internet connection"}
	}

	defer resp.Body.Close()

	if debug {
		fmt.Println("Get response Status:>", resp.Status)
		fmt.Println("Get response Headers:>", resp.Header)
	}

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &restError{err, url, "Error with read Body"}
	}
	if debug {
		fmt.Printf("Body : \n %s \n\n", body)
	}

	if body == nil {
		return &restError{err, url, "Error the body is null, error in the secret key in the config.json ? "}
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body

	return nil
}

// Post Rest on the Nest API
func (r *restHTTP) PostJSON(url string, buffer []byte) (err error) {

	if debug {
		fmt.Println("\n")
		fmt.Println("URL Post :>", url)
		fmt.Printf("Decode Post :> %s \n\n", buffer)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return &restError{err, url, "Check your internet connection"}
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &restError{err, url, "Error with read Body"}
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

	return nil
}

func (r *restHTTP) GetBody() ([]byte, error) {
	return r.body, nil
}
