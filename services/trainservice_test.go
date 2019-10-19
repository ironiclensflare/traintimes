package services

import "testing"

func TestEmptyInputStationReturnsError(t *testing.T) {
	service := NewRealTrainService()
	_, err := service.GetTrainsDepartingFrom("")

	if err == nil {
		t.Error("Did not receive error")
	}
}

func TestTooShortInputStationReturnsError(t *testing.T) {
	service := NewRealTrainService()
	_, err := service.GetTrainsDepartingFrom("AA")

	if err == nil {
		t.Error("Did not receive error")
	}
}

func TestTooLongInputStationReturnsError(t *testing.T) {
	service := NewRealTrainService()
	_, err := service.GetTrainsDepartingFrom("AAAA")

	if err == nil {
		t.Error("Did not receive error")
	}
}
