package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mylog"
	"net/http"
)

type restHTTP struct {
	status string
	header http.Header
	body   []byte
}

type RestHTTP interface {
	GetBody() []byte
	Get(url string) (err error)
	PostJSON(url string, buffer []byte) (err error)
}

type restError struct {
	message error
	url     string
	advice  string
}

func (e *restError) Error() string {
	return fmt.Sprintf("\n \t RestError :> %s \n\t Rest URL :> %s \n\t Rest Advice :> %s", e.message, e.url, e.advice)
}

func MakeNew() (rest RestHTTP) {
	return &restHTTP{}
}

// Get Rest on the Nest API
func (r *restHTTP) Get(url string) (err error) {

	mylog.Trace.Println("Rest Get URL:>", url)

	resp, err := http.Get(url)
	if err != nil {
		return &restError{err, url, "Check your internet connection or if the site is alive"}
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &restError{err, url, "Error with read Body"}
	}

	mylog.Trace.Printf("Body : \n %s \n\n", body)

	if body == nil {
		return &restError{err, url, "Error the body is null, error in the secret key in the config.json ? "}
	}

	if resp.StatusCode != 200 {
		fmt.Println("\n URL Get :>", url)
		fmt.Println("Get response Status:>", resp.Status)
		fmt.Println("Get response Headers:>", resp.Header)
		fmt.Println("Get response Body:>", string(body))
	}

	mylog.Trace.Println("\n URL Get :>", url)
	mylog.Trace.Println("Get response Status:>", resp.Status)
	mylog.Trace.Println("Get response Headers:>", resp.Header)
	mylog.Trace.Println("Get response Body:>", string(body))

	if resp.StatusCode != 200 {
		return &restError{err, url, "Error Status Post"}
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body

	return nil
}

// Post Rest on the Nest API
func (r *restHTTP) PostJSON(url string, buffer []byte) (err error) {

	mylog.Trace.Println("\n")
	mylog.Trace.Println("URL Post :>", url)
	mylog.Trace.Println("Decode Post :> %s \n\n", buffer)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return &restError{err, url, "Check your internet connection or if the site is alive"}
	}

	defer resp.Body.Close()

	//read Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &restError{err, url, "Error with read Body"}
	}

	if body == nil {
		fmt.Println("Error the body is null, error in the secret key in the config.json ? ")
		mylog.Error.Fatal(err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println("\n URL Post :>", url)
		fmt.Printf("Decode Post :> %s \n\n", buffer)
		fmt.Println("Post response Status:>", resp.Status)
		fmt.Println("Post response Headers:>", resp.Header)
		fmt.Println("Post response Body:>", string(body))
	}

	mylog.Trace.Println("\n URL Post :>", url)
	mylog.Trace.Printf("Decode Post :> %s \n\n", buffer)
	mylog.Trace.Println("Post response Status:>", resp.Status)
	mylog.Trace.Println("Post response Headers:>", resp.Header)
	mylog.Trace.Println("Post response Body:>", string(body))

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		fmt.Println("Post response Status:>", resp.Status)
		return &restError{err, url, "Error Status Post"}
	}

	r.status = resp.Status
	r.header = resp.Header
	r.body = body

	return nil
}

func (r *restHTTP) GetBody() []byte {
	return r.body
}
