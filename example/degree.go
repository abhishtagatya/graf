package main

import (
	"fmt"
	"graf"
	"sort"
)

func main() {
	graph, err := graf.FromFile("./data/dimacs/NY/USA-road-d.NY.gr", "a")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inDegree := make(map[string]int)
	outDegree := make(map[string]int)

	for v, edge := range graph.Edges {
		if _, ok := outDegree[v.Id]; !ok {
			outDegree[v.Id] = len(edge)
		}

		if _, ok := inDegree[v.Id]; !ok {
			inDegree[v.Id] = 0
		}

		for _, ev := range edge {
			if _, ok := inDegree[ev.ConnectedId]; !ok {
				inDegree[ev.ConnectedId] = 0
			}

			inDegree[ev.ConnectedId] += 1
		}
	}

	for v := range inDegree {
		fmt.Println(v, inDegree[v], outDegree[v])
	}

	// Extract keys into a slice
	values := make([]int, 0, len(outDegree))
	for _, v := range outDegree {
		values = append(values, v)
	}

	// Sort the values
	sort.Ints(values)

	fmt.Println(values)

}
