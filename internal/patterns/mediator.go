/* Mediator is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly,
through a special mediator object. */

/* An excellent example of the Mediator pattern is a railway station traffic system.
Two trains never communicate between themselves for the availability of the platform.
The stationManager acts as a mediator and makes the platform available to only one of the arriving trains while keeping the rest in a queue.
A departing train notifies the stations, which lets the next train in the queue to arrive. */

package patterns

import "fmt"

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

// train types

type PassengerTrain struct {
	Mediator Mediator
	Id       string
}

func (p *PassengerTrain) Arrive() { // controls to arrive
	if !p.Mediator.canArrive(p) {
		fmt.Printf("%s -> Arrive blocked waiting... \n ", p.Id)
		return
	}

	fmt.Printf("%s -> Arrived \n ", p.Id)
}

func (p *PassengerTrain) Depart() { // controls to train movement
	p.Mediator.notifyAboutDeparture()
	fmt.Printf("%s -> Leaving \n ", p.Id)
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Printf("%s -> Arrival permitted, arriving \n ", p.Id)
	p.Arrive()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
