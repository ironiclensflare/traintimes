package main

import "github.com/DATA-DOG/godog"

func iHaveATrainClientWithRealAPICredentials() error {
	return godog.ErrPending
}

func iHaveEnteredAValidStation() error {
	return godog.ErrPending
}

func iGetTheDepartureBoard() error {
	return godog.ErrPending
}

func iShouldReceiveAResponse() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a Train Client with real API credentials$`, iHaveATrainClientWithRealAPICredentials)
	s.Step(`^I have entered a valid station$`, iHaveEnteredAValidStation)
	s.Step(`^I get the departure board$`, iGetTheDepartureBoard)
	s.Step(`^I should receive a response$`, iShouldReceiveAResponse)
}
