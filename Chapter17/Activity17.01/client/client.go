package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Names struct {
	Names []string `json:"names"`
}

func getDataAndParseResponse() (int, int) {

	// send the GET request
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.Body.Close()

	names := Names{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}

	electricCount := 0
	boogalooCount := 0
	for _, name := range names.Names {
		if name == "Electric" {
			electricCount++
		} else if name == "Boogaloo" {
			boogalooCount++
		}
	}

	// return the count data
	return electricCount, boogalooCount
}

func main() {
	electricCount, boogalooCount := getDataAndParseResponse()
	fmt.Println("Electric Count: ", electricCount)
	fmt.Println("Boogaloo Count: ", boogalooCount)
}
