package main

import (
	"fmt"
	"log"
	"sotoon/dijkstra"
	"sotoon/model"
	"time"
)

func main() {
	data := model.NewRoads("roads.json")

	graph := dijkstra.NewGraph(data)

	travel := model.NewTravel("input.json")

	t, err := time.Parse(time.Kitchen, travel.DepartureTime)
	if err != nil {
		log.Fatal(err)
	}

	costTable := graph.Dijkstra(travel.Source, t)
	bestPath := graph.ShortestPath(costTable, travel.Destination, travel.Passengers)
	fmt.Println(bestPath)
}
