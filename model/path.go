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
	return fmt.Sprintf("%s(%s)-->%s(%s)\nduration %f minutes", e.Source, e.StartTime.Format(time.Kitchen), e.Destination, e.ArrivalTime.Format(time.Kitchen), e.ArrivalTime.Sub(e.StartTime).Minutes())
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
}

func (s ShortestPath) String() (path string) {
	for _, p := range s.BestPath {
		path = p.String() + path
	}

	path += fmt.Sprintf("Total duration: %f minutes", s.ArrivalTime.Sub(s.StartTime).Minutes())

	return path
}
