package services

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const URL = "https://api.rtt.io/api/v1/json/search/"

type TrainService interface {
	GetTrainsDepartingFrom(station string) (json string, err error)
}

type RealTrainService struct {
}

func NewRealTrainService() *RealTrainService {
	return &RealTrainService{}
}

func (service *RealTrainService) GetTrainsDepartingFrom(station string) (json string, err error) {
	if len(station) != 3 {
		return "", errors.New("Invalid station name")
	}
	req, err := http.NewRequest("GET", URL+station, nil)
	req.SetBasicAuth(os.Getenv("RTTUSER"), os.Getenv("RTTPASS"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
