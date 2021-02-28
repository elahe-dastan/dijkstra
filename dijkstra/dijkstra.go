package dijkstra

import (
	"fmt"
	"sort"
	"sotoon/model"
	"time"
)

type Graph struct {
	Edges     []*Edge
	Nodes     []*Node
	Cities    map[string]*Node
}

func NewGraph(data model.Roads) *Graph {
	graph := Graph{Cities: make(map[string]*Node)}

	for _, road := range data.RoadsArr {
		graph.Cities[road.Source] = &Node{Name: road.Source}
		graph.Cities[road.Destination] = &Node{Name: road.Destination}
	}

	for _, road := range data.RoadsArr {
		graph.AddEdge(graph.Cities[road.Source], graph.Cities[road.Destination], float64(road.Length)/float64(road.SpeedLimit), road.Length, road.RoadName)
	}

	return &graph
}

type Edge struct {
	Parent   *Node
	Child    *Node
	Cost     float64
	Length   int
	RoadName string
}

type Node struct {
	Name string
}

type CostPrevious struct {
	Cost     float64
	Previous *Node
	Path     string
	Arrival  time.Time
}

type CostTable struct {
	StartTime time.Time
	StartNode *Node
	Table     map[*Node]*CostPrevious
}

const Infinity = float64(^uint(0) >> 1)

// AddEdge adds an Edge to the Graph
func (g *Graph) AddEdge(parent, child *Node, cost float64, length int, roadName string) {
	edge := &Edge{
		Parent:   parent,
		Child:    child,
		Cost:     cost,
		Length:   length,
		RoadName: roadName,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

// AddNode adds a Node to the Graph list of Nodes, if the the node wasn't already added
func (g *Graph) AddNode(node *Node) {
	var isPresent bool
	for _, n := range g.Nodes {
		if n == node {
			isPresent = true
		}
	}

	if !isPresent {
		g.Nodes = append(g.Nodes, node)
	}
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += fmt.Sprintf("%s -> %s = %f", edge.Parent.Name, edge.Child.Name, edge.Cost)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}
	s += "\n"

	return s
}

// Dijkstra implements THE Dijkstra algorithm
// Returns the shortest path from startNode to all the other Nodes
func (g *Graph) Dijkstra(startCity string, startTime time.Time) *CostTable {
	startNode := g.Cities[startCity]
	// First, we instantiate a "Cost Table", it will hold the information:
	// "From startNode, what's is the cost to all the other Nodes?"
	// When initialized, It looks like this:
	// NODE  COST
	//  A     0    // The startNode has always the lowest cost to itself, in this case, 0
	//  B    Inf   // the distance to all the other Nodes are unknown, so we mark as Infinity
	//  C    Inf
	// ...
	costTable := g.NewCostTable(startNode, startTime)

	// An empty list of "visited" Nodes. Everytime the algorithm runs on a Node, we add it here
	var visited []*Node

	// A loop to visit all Nodes
	for len(visited) != len(g.Nodes) {

		// Get closest non visited Node (lower cost) from the costTable
		node := getClosestNonVisitedNode(costTable.Table, visited)

		// Mark Node as visited
		visited = append(visited, node)

		// Get Node's Edges (its neighbors)
		nodeEdges := g.GetNodeEdges(node)

		for _, edge := range nodeEdges {

			// The distance to that neighbor, let's say B is the cost from the costTable + the cost to get there (Edge cost)
			// In the first run, the costTable says it's "Infinity"
			// Plus the actual cost, let's say "5"
			// The distance becomes "5"
			table := costTable.Table
			distanceToNeighbor := table[node].Cost + edge.Cost

			// If the distance above is lesser than the distance currently in the costTable for that neighbor
			if distanceToNeighbor < table[edge.Child].Cost {

				// Update the costTable for that neighbor
				table[edge.Child].Cost = distanceToNeighbor
				table[edge.Child].Previous = edge.Parent
				table[edge.Child].Path = edge.RoadName
				table[edge.Child].Arrival = table[edge.Child].Arrival.Add(time.Duration(distanceToNeighbor*3600) * time.Second)
			}
		}
	}

	//Make the costTable nice to read :)
	//shortestPathTable := ""
	//for node, cost := range costTable {
	//	shortestPathTable += fmt.Sprintf("Distance from %s to %s = %d previous = %s\n", startNode.Name, node.Name, cost.Cost, cost.Previous.Name)
	//}

	//fmt.Println(shortestPathTable)

	return costTable
}

// NewCostTable returns an initialized cost table for the Dijkstra algorithm work with
// by default, the lowest cost is assigned to the startNode – so the algorithm starts from there
// all the other Nodes in the Graph receives the Infinity value
func (g *Graph) NewCostTable(startNode *Node, startTime time.Time) *CostTable {
	costTable := &CostTable{
		StartTime: startTime,
		StartNode: startNode,
		Table:     make(map[*Node]*CostPrevious),
	}

	table := costTable.Table
	table[startNode] = &CostPrevious{Arrival: startTime}

	for _, node := range g.Nodes {
		if node != startNode {
			table[node] = &CostPrevious{
				Cost:    Infinity,
				Arrival: startTime,
			}
		}
	}

	return costTable
}

// GetNodeEdges returns all the Edges that start with the specified Node
// In other terms, returns all the Edges connecting to the Node's neighbors
func (g *Graph) GetNodeEdges(node *Node) (edges []*Edge) {
	for _, edge := range g.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

// getClosestNonVisitedNode returns the closest Node (with the lower cost) from the costTable
// **if the node hasn't been visited yet**
func getClosestNonVisitedNode(costTable map[*Node]*CostPrevious, visited []*Node) *Node {
	type CostTableToSort struct {
		Node *Node
		Cost float64
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, CostTableToSort{node, cost.Cost})
		}
	}

	// We need the Node with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

func (g *Graph) ShortestPath(costTable *CostTable, destCity string) model.ShortestPath {
	dest := g.Cities[destCity]

	s := model.ShortestPath{
		BestPath:    make([]model.Path, 0),
		StartTime:   costTable.StartTime,
		ArrivalTime: costTable.Table[dest].Arrival,
	}

	for dest != costTable.StartNode {
		p := model.Path{
			Reference: costTable.Table[dest].Path,
			Route: model.Edge{
				Source:      costTable.Table[dest].Previous.Name,
				StartTime:   costTable.Table[costTable.Table[dest].Previous].Arrival,
				Destination: dest.Name,
				ArrivalTime: costTable.Table[dest].Arrival,
			},
		}

		dest = costTable.Table[dest].Previous

		s.BestPath = append(s.BestPath, p)
	}

	return s
}
