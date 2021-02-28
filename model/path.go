package model

import (
	"fmt"
	"time"
)

type Edge struct {
	Source      string
	StartTime   time.Time
	Destination string
	ArrivalTime time.Time
	Duration    time.Duration
	Children    int
	Adults      int
	Cars        int
	Price       int
}

func (e Edge) String() string {
	return fmt.Sprintf("%s(%s)-->%s(%s)\nduration %f minutes\n children(%d) adults(%d)\n%d car(s) required\ntotal: %d",
		e.Source, e.StartTime.Format(time.Kitchen), e.Destination, e.ArrivalTime.Format(time.Kitchen),
		e.ArrivalTime.Sub(e.StartTime).Minutes(), e.Children, e.Adults, e.Cars, e.Price)
}

type Path struct {
	Type      string
	Reference string
	Route     Edge
}

func (p Path) String() string {
	return fmt.Sprintf("type: Road\nrefrence: %s\nroute:%s\n********************\n", p.Reference, p.Route)
}

type ShortestPath struct {
	BestPath    []Path
	StartTime   time.Time
	ArrivalTime time.Time
	Passengers  []int
}

func (s ShortestPath) String() (path string) {
	totalPrice := 0
	for _, p := range s.BestPath {
		path = p.String() + path
		totalPrice += p.Route.Price
	}

	path += fmt.Sprintf("Total duration: %f minutes\n", s.ArrivalTime.Sub(s.StartTime).Minutes())
	path += fmt.Sprintf("Total price: %d", totalPrice)

	return path
}
