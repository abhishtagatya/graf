package graf

import (
	"errors"
	"fmt"
	"math"
)

func BreadthSearch(graph Graph, s string) (*AlgorithmReport, error) {
	var sv Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}

	report := AlgorithmReport{
		StartVertex:    &sv,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := BlankQueue()
	queue.Push(&QueueItem{Value: sv})

	for !queue.IsEmpty() {
		cq := queue.Pop().(*QueueItem)
		cv := cq.Value.(Vertex)

		report.EndVertex = &cv
		report.Distance = report.DistanceMap[cv.Id]
		report.PredecessorChain = append(report.PredecessorChain, cv)

		report.VisitMap[cv.Id] = true
		for _, edge := range graph.Edges[cv] {
			if ok = report.VisitMap[edge.ConnectedId]; !ok {
				report.DistanceMap[edge.ConnectedId] = report.DistanceMap[cv.Id] + 1
				report.PredecessorMap[edge.ConnectedId] = &cv

				queue.Push(&QueueItem{Value: *edge.ConnectedVertex})
			}
		}
	}

	return &report, nil
}
