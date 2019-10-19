package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const URL = "https://api.rtt.io/api/v1/json/search/"

type TrainService interface {
	GetTrainsDepartingFrom(station string) (json string)
}

type RealTrainService struct {
}

type FakeTrainService struct {
}

func (service *FakeTrainService) GetTrainsDepartingFrom(station string) (json string) {
	return `{}`
}

func NewRealTrainService() *RealTrainService {
	return &RealTrainService{}
}

func NewFakeTrainService() *FakeTrainService {
	return &FakeTrainService{}
}

func (service *RealTrainService) GetTrainsDepartingFrom(station string) (json string) {
	req, err := http.NewRequest("GET", URL+station, nil)
	req.SetBasicAuth(os.Getenv("RTTUSER"), os.Getenv("RTTPASS"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
