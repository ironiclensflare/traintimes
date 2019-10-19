package models

type DepartureBoard struct {
	Location Location
}

type Location struct {
	Name   string
	CRS    string
	TIPLOC string
}
