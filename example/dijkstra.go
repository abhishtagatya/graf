package main

import (
	"fmt"
	"graf"
	"strconv"
)

func main() {
	graph, err := graf.FromFile("./data/dimacs/NY/USA-road-d.NY.gr", "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := strconv.FormatInt(int64(6), 10)
	e := strconv.FormatInt(int64(1), 10)

	if s == e {
		return
	}

	result, err := graf.Dijkstra(*graph, s, e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result.Distance, len(result.PredecessorChain), len(graph.Vertices))
}