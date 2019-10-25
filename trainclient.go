package main

import (
	"encoding/json"
	"os"

	"github.com/ironiclensflare/traintimes/models"
	"github.com/ironiclensflare/traintimes/services"
)

// TrainClient represents a train client
type TrainClient struct {
	username string
	password string
}

// GetDepartureBoard retrieves a departure board for the given station
func (tc TrainClient) GetDepartureBoard(station string) models.DepartureBoard {
	var service services.TrainService
	service = services.NewRealTrainService()
	trainsJSON, _ := service.GetTrainsDepartingFrom(station)
	var trains models.DepartureBoard
	json.Unmarshal([]byte(trainsJSON), &trains)
	return trains
}

// GetTrainClient returns a new instance of TrainClient with credentials from the environment variables
func GetTrainClient() *TrainClient {
	client := TrainClient{username: os.Getenv("RTTUSER"), password: os.Getenv("RTTPASS")}
	return &client
}

// GetTrainClientWithCredentials returns a new instance of TrainClient with the given credentials
func GetTrainClientWithCredentials(username string, password string) *TrainClient {
	client := TrainClient{username: username, password: password}
	return &client
}
