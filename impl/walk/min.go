package walk

import (
	"errors"
	"fmt"
	"graf"
	"math"
)

// MinWalk Traverses the graph through the minimum edge weight in a specified number of step
func MinWalk(graph graf.Graph, s string, step int) (*graf.AlgorithmReport, error) {
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
		var mv graf.Vertex
		minWeight := math.Inf(1)

		for _, edge := range graph.Edges[sv] {
			if edge.Weight < minWeight {
				minWeight = edge.Weight
				mv = *edge.ConnectedVertex
			}
		}

		report.EndVertex = &mv
		report.VisitMap[mv.Id] = true
		report.Distance = report.Distance + minWeight
		report.DistanceMap[mv.Id] = report.Distance
		report.PredecessorMap[mv.Id] = &sv

		sv = mv
	}

	return &report, nil
}
