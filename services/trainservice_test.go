package services

import "testing"

func TestCorrectLengthStationDoesNotReturnError(t *testing.T) {
	service := NewRealTrainService()
	_, err := service.GetTrainsDepartingFrom("NOT")

	if err != nil {
		t.Error("Should not have received error")
	}
}

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
