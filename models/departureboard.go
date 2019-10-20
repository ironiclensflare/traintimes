package models

// DepartureBoard represents a departure board
type DepartureBoard struct {
	Location Location
}

// Location represents a location
type Location struct {
	Name   string
	CRS    string
	TIPLOC string
}
