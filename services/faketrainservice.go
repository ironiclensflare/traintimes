package services

type FakeTrainService struct{}

func (service *FakeTrainService) GetTrainsDepartingFrom(station string) (json string, err error) {
	return `{}`, nil
}

func NewFakeTrainService() *FakeTrainService {
	return &FakeTrainService{}
}
