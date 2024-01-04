package graf

import (
	"errors"
	"fmt"
	"math"
)

func Breadth(graph Graph, s string) (*AlgorithmReport, error) {
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

func Depth(graph Graph) (*AlgorithmReport, error) {
	var ok bool

	report := AlgorithmReport{
		Distance:       0,
		DistanceMap:    map[string]float64{},
		PredecessorMap: map[string]*Vertex{},
		VisitMap:       map[string]bool{},
	}

	for _, v := range graph.Vertices {
		if ok = report.VisitMap[v.Id]; !ok {
			depthVisit(graph, v, &report)
		}
	}

	return &report, nil
}

func depthVisit(graph Graph, u Vertex, report *AlgorithmReport) {
	var ok bool

	// Distance as Discovery
	report.VisitMap[u.Id] = true
	report.Distance += 1
	report.DistanceMap[u.Id] = report.Distance
	report.PredecessorChain = append(report.PredecessorChain, u)

	for _, edge := range graph.Edges[u] {
		if report.StartVertex == nil {
			report.StartVertex = &u
		}

		if ok = report.VisitMap[edge.ConnectedId]; !ok {
			depthVisit(graph, *edge.ConnectedVertex, report)
		}
	}

	report.EndVertex = &u
}
