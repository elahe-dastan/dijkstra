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

	t, err := time.Parse(time.Kitchen, "8:00AM")
	if err != nil {
		log.Fatal(err)
	}

	graph := dijkstra.NewGraph(data)

	costTable := graph.Dijkstra("TEHRAN", t)
	bestPath := graph.ShortestPath(costTable, "RAMSAR")
	fmt.Println(bestPath)
}
