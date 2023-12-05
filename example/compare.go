package main

import (
	"fmt"
	"graf"
	"graf/auxiliary"
	"strconv"
)

func main() {
	fmt.Println("Reading Graph")
	graph, err := graf.FromFile("./data/dimacs/BAY/USA-road-d.BAY.50k.gr", "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Reading Aux")
	auxContainer, err := auxiliary.LoadContainer("./data/dimacs/BAY/USA-road-d.BAY.50k.gr.aux", "x")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := strconv.FormatInt(int64(5186), 10)
	e := strconv.FormatInt(int64(9917), 10)

	if s == e {
		return
	}

	fmt.Println("Running Dijkstra")
	result, err := graf.Dijkstra(*graph, s, e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result.Distance, len(result.VisitMap), len(graph.Vertices))

	fmt.Println("Running Dijkstra (Pruning)")
	resultT, err := auxiliary.DijkstraGeometricPrune(*graph, s, e, auxContainer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resultT.Distance, len(resultT.VisitMap), len(graph.Vertices))
}
