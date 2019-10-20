package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ironiclensflare/traintimes/models"
	"github.com/ironiclensflare/traintimes/services"
)

func main() {
	getTrains()
}

func getTrains() {
	args := os.Args[1:]
	var service services.TrainService
	service = services.NewRealTrainService()
	trainsJSON, _ := service.GetTrainsDepartingFrom(args[0])
	var trains models.DepartureBoard
	json.Unmarshal([]byte(trainsJSON), &trains)

	for _, s := range trains.Services {
		fmt.Printf("%v - %v (%v)\n", s.LocationDetail.RealtimeDeparture, s.LocationDetail.Destination[0].Description, s.ATOCName)
	}
}
