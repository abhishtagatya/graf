package minoverlay

import (
	"container/heap"
	"errors"
	"fmt"
	"github.com/abhishtagatya/graf"
)

type OverlayPair struct {
	Weight float64
	Sigma  float64
}

func DijkstraOverlay(graph graf.Graph, sv graf.Vertex, sub []string) map[string]OverlayPair {
	distances := make(map[string]OverlayPair)
	visits := make(map[string]bool)

	distances[sv.Id] = OverlayPair{Weight: 0, Sigma: 0}
	visits[sv.Id] = false

	queue := graf.BlankMinPriorityQueue()
	heap.Push(&queue, &graf.QueueItem{
		Value:    sv,
		Priority: 0,
	})

	subset := make([]string, len(sub))
	copy(subset, sub)

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*graf.QueueItem)
		cv := cq.Value.(graf.Vertex)

		if graf.ContainsVertex(subset, cv.Id) {
			subset = removeElementByValue(subset, cv.Id)
		}

		if len(subset) == 0 {
			return distances
		}

		if visits[cv.Id] {
			continue
		}

		visits[cv.Id] = true
		for _, edge := range graph.Edges[cv] {
			var newPair OverlayPair

			if graf.ContainsVertex(subset, cv.Id) && !graf.ContainsVertex(subset, edge.ConnectedId) {
				newPair = OverlayPair{Weight: distances[cv.Id].Weight + edge.Weight, Sigma: -1}
			} else {
				newPair = OverlayPair{Weight: distances[cv.Id].Weight + edge.Weight, Sigma: 0}
			}

			if neighPair, ok := distances[edge.ConnectedId]; !ok || newPair.Weight < neighPair.Weight {
				distances[edge.ConnectedId] = newPair
				heap.Push(&queue, &graf.QueueItem{
					Value:    *edge.ConnectedVertex,
					Priority: newPair.Weight,
				})
			}
		}
	}
	return distances
}

func MinOverlay(graph graf.Graph, l float64, sub []string) (*OverlayGraph, error) {
	aux := BlankOverlayGraph()

	for _, u := range sub {
		uv, ok := graph.Vertices[u]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Vertex: %s is not in Graph.", u))
		}

		distances := DijkstraOverlay(graph, uv, sub)
		for _, v := range sub {
			ev, ok := graph.Vertices[v]
			if !ok {
				return nil, errors.New(fmt.Sprintf("Vertex: %s is not in Graph.", v))
			}

			if v == u {
				continue
			}

			if distances[v].Sigma == 0 {
				aux.Vertices[v] = uv
				aux.AddEdge(uv, ev, distances[v].Weight)
				aux.AddOverlayEntry(uv, ev, distances[v])
			}
		}
	}

	return &aux, nil
}
