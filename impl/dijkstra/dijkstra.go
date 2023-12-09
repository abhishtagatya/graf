package dijkstra

import (
	"container/heap"
	"errors"
	"fmt"
	"graf"
	"math"
)

// SingleDijkstra Single-Source Single-Target Dijkstra
func SingleDijkstra(graph graf.Graph, s string, e string) (*graf.AlgorithmReport, error) {
	var sv graf.Vertex
	var ev graf.Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}
	if ev, ok = graph.Vertices[e]; !ok {
		return nil, errors.New(fmt.Sprintf("Ending Vertex: %s is not in Graph.", e))
	}

	report := graf.AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      &ev,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*graf.Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := graf.BlankQueue()
	heap.Push(&queue, &graf.QueueItem{
		Value:    sv,
		Priority: 0,
	})

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*graf.QueueItem)
		cv := cq.Value.(graf.Vertex)

		if cv == ev {
			report.Distance = report.DistanceMap[ev.Id]

			pv := &ev
			for pv != nil {
				report.PredecessorChain = append(report.PredecessorChain, *pv)
				tv := report.PredecessorMap[pv.Id]
				pv = tv
			}

			return &report, nil
		}

		if report.VisitMap[cv.Id] {
			continue
		}

		report.VisitMap[cv.Id] = true
		for _, edge := range graph.Edges[cv] {
			newDist := cq.Priority + edge.Weight
			if dist, ok := report.DistanceMap[edge.ConnectedId]; !ok || newDist < dist {
				report.DistanceMap[edge.ConnectedId] = newDist
				report.PredecessorMap[edge.ConnectedId] = &cv
				heap.Push(&queue, &graf.QueueItem{
					Value:    *edge.ConnectedVertex,
					Priority: newDist,
				})
			}
		}
	}

	return &report, nil
}
