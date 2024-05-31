package main

import (
	patterns "golang-design-patterns/internal/patterns"
)

func main() {
	stationManager := patterns.NewStationManager()

	passengerTrain1 := &patterns.PassengerTrain{
		Mediator: stationManager,
		Id:       "PassengerTrain1",
	}

	passengerTrain2 := &patterns.PassengerTrain{
		Mediator: stationManager,
		Id:       "PassengerTrain2",
	}

	passengerTrain1.Arrive()
	passengerTrain2.Arrive()
	passengerTrain1.Depart()

}
