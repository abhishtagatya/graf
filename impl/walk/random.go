package walk

import (
	"errors"
	"fmt"
	"graf"
	"math/rand"
)

// RandomWalk Traverses the graph through a random selection of edges in a specified number of steps
func RandomWalk(graph graf.Graph, s string, step int) (*graf.AlgorithmReport, error) {
	var sv graf.Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}

	report := graf.AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      nil,
		Distance:       0,
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: true},
	}

	for i := 0; i < step; i++ {
		randIdx := rand.Intn(len(graph.Edges[sv]))
		ev := graph.Edges[sv][randIdx]

		report.EndVertex = ev.ConnectedVertex
		report.VisitMap[ev.ConnectedId] = true
		report.Distance = report.Distance + ev.Weight
		report.DistanceMap[ev.ConnectedId] = report.Distance
		report.PredecessorMap[ev.ConnectedId] = &sv

		sv = *ev.ConnectedVertex
	}

	return &report, nil
}
