package main

import (
	"fmt"
	"graf"
	"graf/imp"
)

func main() {
	graph, err := graf.FromFile("./data/dimacs/NY/USA-road-d.NY.gr", "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	coordMap, err := imp.LoadCoordinate("./data/dimacs/NY/USA-road-d.NY.co", "v")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := "30657"
	e := "47091"

	result, err := graf.Dijkstra(*graph, s, e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resultH, err := imp.AStarHaversine(*graph, s, e, coordMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Dijkstra", result.Distance, len(result.VisitMap))
	fmt.Println("Dijkstra (Heuristic)", resultH.Distance, len(resultH.VisitMap))

	fmt.Println(len(result.PredecessorChain))
	fmt.Println(len(resultH.PredecessorChain))
}
