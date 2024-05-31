package patterns

import (
	"testing"
)

func TestStationManager(t *testing.T) {
	stationManager := NewStationManager()

	passengerTrain1 := &PassengerTrain{
		Mediator: stationManager,
		Id:       "PassengerTrain1",
	}
	passengerTrain2 := &PassengerTrain{
		Mediator: stationManager,
		Id:       "PassengerTrain2",
	}

	// Tests

	// Test 1: Arrival of a train when the platform is empty
	passengerTrain1.PermitArrival()
	if len(stationManager.trainQueue) != 0 {
		t.Errorf("Expected: 0, Got: %d", len(stationManager.trainQueue))
	}

	// Test 2: Arrival of a train when the platform is occupied
	passengerTrain2.PermitArrival()
	if len(stationManager.trainQueue) != 1 {
		t.Errorf("Expected: 1, Got: %d", len(stationManager.trainQueue))
	}

	// Test 3: Departure of a train
	passengerTrain1.Depart()
	if len(stationManager.trainQueue) != 0 {
		t.Errorf("Expected: 0, Got: %d", len(stationManager.trainQueue))
	}
}
