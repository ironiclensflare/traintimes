package main

import (
	"encoding/json"
	"os"

	"github.com/olekukonko/tablewriter"

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
	buildTable(trains)
}

func buildTable(trains models.DepartureBoard) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Time", "ID", "Destination", "TOC", "Platform"})
	for _, s := range trains.Services {
		table.Append([]string{s.LocationDetail.RealtimeDeparture, s.TrainIdentity, s.LocationDetail.Destination[0].Description, s.ATOCName, s.LocationDetail.Platform})
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
