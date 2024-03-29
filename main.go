package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/ironiclensflare/traintimes/models"
)

func main() {
	getTrains()
}

func getTrains() {
	args := os.Args[1:]
	client := GetTrainClient()
	trains := client.GetDepartureBoard(args[0])
	buildTable(trains)
}

func buildTable(trains models.DepartureBoard) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, fmt.Sprintf("Trains at %s", trains.Location.Name))
	table.SetHeader([]string{"Time", "Head", "Destination", "TOC", "Platform"})
	for _, s := range trains.Services {
		var platform string
		if s.LocationDetail.IsCancelled() {
			red := color.New(color.BgRed).SprintFunc()
			platform = red("CANCELLED")
		} else {
			platform = s.LocationDetail.Platform
		}
		table.Append([]string{s.LocationDetail.RealtimeDeparture, s.TrainIdentity, s.LocationDetail.Destination[0].Description, s.ATOCName, platform})
	}
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
