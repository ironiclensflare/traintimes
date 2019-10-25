package main

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/godog"
	"github.com/ironiclensflare/traintimes/models"
)

var client *TrainClient
var stationToQuery string
var response models.DepartureBoard

func iHaveATrainClientWithRealAPICredentials() error {
	username := os.Getenv("RTTUSER")
	password := os.Getenv("RTTPASS")
	client = GetTrainClientWithCredentials(username, password)
	return nil
}

func iHaveEnteredAValidStation() error {
	stationToQuery = "VIC"
	return nil
}

func iGetTheDepartureBoard() error {
	response = client.GetDepartureBoard(stationToQuery)
	return nil
}

func iShouldReceiveAResponse() error {
	if response.Location.Name == "" {
		return fmt.Errorf("Didn't receive any services")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a Train Client with real API credentials$`, iHaveATrainClientWithRealAPICredentials)
	s.Step(`^I have entered a valid station$`, iHaveEnteredAValidStation)
	s.Step(`^I get the departure board$`, iGetTheDepartureBoard)
	s.Step(`^I should receive a response$`, iShouldReceiveAResponse)
}
