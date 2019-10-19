package main

import (
	"log"

	"github.com/ironiclensflare/traintimes/services"
)

func main() {
	getTrains()
}

func getTrains() {
	var service services.TrainService
	service = services.NewRealTrainService()
	json := service.GetTrainsDepartingFrom("NOT")
	log.Println(json)
}
