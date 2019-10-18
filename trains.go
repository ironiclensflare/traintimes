package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Getting trains...")
	getTrains()
}

const URL = "https://api.rtt.io/api/v1/json/search/BMH"

func getTrains() {
	req, err := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(os.Getenv("RTTUSER"), os.Getenv("RTTPASS"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
