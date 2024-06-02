package main

import (
	patterns "golang-design-patterns/internal/patterns"
)

func mediator() {
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

func command() {
	tv := &patterns.Tv{}

	onCommand := &patterns.OnCommand{
		Device: tv,
	}

	offCommand := &patterns.OffCommand{
		Device: tv,
	}

	onButton := &patterns.Button{
		Command: onCommand,
	}
	onButton.Press()

	offButton := &patterns.Button{
		Command: offCommand,
	}
	offButton.Press()
}

func main() {
	mediator()
	command()
}
