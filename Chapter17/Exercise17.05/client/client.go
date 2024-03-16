package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse() string {
	client := http.Client{Timeout: 11 * time.Second}

	// create the GET request
	req, err := http.NewRequest("POST", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "superSecretToken")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// return the response data
	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
