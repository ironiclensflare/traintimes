package models

// DepartureBoard represents a departure board
type DepartureBoard struct {
	Location Location
	Services []Service
}

// Location represents a location
type Location struct {
	Name   string
	CRS    string
	TIPLOC string
}

// Service represents a train service
type Service struct {
	LocationDetail  LocationDetail
	ServiceUID      string
	RunDate         string
	TrainIdentity   string
	RunningIdentity string
	ATOCCode        string
	ATOCName        string
	ServiceType     string
	IsPassenger     bool
	Origin          []Pair
	Destination     []Pair
}

// LocationDetail represents a location
type LocationDetail struct {
	Origin                  []Pair
	Destination             []Pair
	IsCall                  bool
	IsPublicCall            bool
	RealtimeArrival         string
	RealtimeArrivalActual   bool
	RealtimeDeparture       string
	RealtimeDepartureActual bool
	Platform                string
	PlatformConfirmed       bool
	PlatformChanged         bool
	DisplayAs               bool
}

// Pair represents an origin or destination station
type Pair struct {
	TIPLOC      string
	Description string
	WorkingTime string
	PublicTime  string
}
