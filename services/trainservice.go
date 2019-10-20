package services

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "https://api.rtt.io/api/v1/json/search/"

// TrainService is an interface for getting trains
type TrainService interface {
	GetTrainsDepartingFrom(station string) (json string, err error)
}

// RealTrainService is a real train service
type RealTrainService struct {
}

// NewRealTrainService gets an instance of RealTrainService
func NewRealTrainService() *RealTrainService {
	return &RealTrainService{}
}

// GetTrainsDepartingFrom gets a list of trains departing from the given station
func (service *RealTrainService) GetTrainsDepartingFrom(station string) (json string, err error) {
	if len(station) != 3 {
		return "", errors.New("Invalid station name")
	}
	req, err := http.NewRequest("GET", url+station, nil)
	req.SetBasicAuth(os.Getenv("RTTUSER"), os.Getenv("RTTPASS"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
