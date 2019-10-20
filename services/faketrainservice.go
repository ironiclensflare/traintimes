package services

// FakeTrainService is a fake train service
type FakeTrainService struct{}

// GetTrainsDepartingFrom gets a list of trains departing from the given station
func (service *FakeTrainService) GetTrainsDepartingFrom(station string) (json string, err error) {
	return `{}`, nil
}

// NewFakeTrainService returns an instance of FakeTrainService
func NewFakeTrainService() *FakeTrainService {
	return &FakeTrainService{}
}
