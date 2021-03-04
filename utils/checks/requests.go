package checks

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

// MakeRequest will do the GET request
func MakeRequest(host string) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get(host)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

// MakePostRequest will do the POST request
func MakePostRequest(host string, data []byte, api string) (*http.Request, *http.Response) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	if host == "https://fcm.googleapis.com/fcm/send" {
		request, _ := http.NewRequest("POST", host, bytes.NewBuffer(data))
		request.Header.Set("Authorization", "key="+api)
		request.Header.Set("Content-Type", "application/json")
		response, err := client.Do(request)
		CheckErr(err)
		return request, response
	} else {
		request, _ := http.NewRequest("POST", host+api, bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")
		response, err := client.Do(request)
		CheckErr(err)
		return request, response
	}

}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
