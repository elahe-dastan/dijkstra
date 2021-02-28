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
}

func (e Edge) String() string {
	return fmt.Sprintf("%s(%s)-->%s(%s)\nduration %f minutes",
		e.Source, e.StartTime.Format(time.Kitchen), e.Destination, e.ArrivalTime.Format(time.Kitchen),
		e.ArrivalTime.Sub(e.StartTime).Minutes())
}

type Path struct {
	Type      string
	Reference string
	Route     Edge
	Children    int
	Adults      int
	Cars        int
	Price       int
}

func (p Path) String() string {
	return fmt.Sprintf("type: Road\nrefrence: %s\nroute:%s\nprice: children(%d) adults(%d)\n%d car(s) required\ntotal: %d\n********************\n",
		p.Reference, p.Route, p.Children, p.Adults, p.Cars, p.Price)
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
		totalPrice += p.Price
	}

	path += fmt.Sprintf("Total duration: %f minutes\n", s.ArrivalTime.Sub(s.StartTime).Minutes())
	path += fmt.Sprintf("Total price: %d", totalPrice)

	return path
}
