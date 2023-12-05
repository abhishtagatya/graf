package graf

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// AlgorithmReport Stores information of algorithm runs
type AlgorithmReport struct {
	StartVertex *Vertex `json:"startVertex"`
	EndVertex   *Vertex `json:"endVertex"`

	Distance float64 `json:"minDistance"`

	DistanceMap    map[string]float64 `json:"distanceMap"`
	PredecessorMap map[string]*Vertex `json:"predecessorMap"`
	VisitMap       map[string]bool    `json:"visitMap"`

	PredecessorChain []Vertex `json:"predecessorChain"`
}

// RandomWalk Traverses the graph through a random selection of edges in a specified number of steps
func RandomWalk(graph Graph, s string, step int) (*AlgorithmReport, error) {
	var sv Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}

	report := AlgorithmReport{
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

// MinWalk Traverses the graph through the minimum edge weight in a specified number of step
func MinWalk(graph Graph, s string, step int) (*AlgorithmReport, error) {
	var sv Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}

	report := AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      nil,
		Distance:       0,
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: true},
	}

	for i := 0; i < step; i++ {
		var mv Vertex
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

// MaxWalk Traverses the graph through the maximum edge weight in a specified number of step
func MaxWalk(graph Graph, s string, step int) (*AlgorithmReport, error) {
	var sv Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}

	report := AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      nil,
		Distance:       0,
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: true},
	}

	for i := 0; i < step; i++ {
		var mv Vertex
		maxWeight := math.Inf(-1)

		for _, edge := range graph.Edges[sv] {
			if edge.Weight > maxWeight {
				maxWeight = edge.Weight
				mv = *edge.ConnectedVertex
			}
		}

		report.EndVertex = &mv
		report.VisitMap[mv.Id] = true
		report.Distance = report.Distance + maxWeight
		report.DistanceMap[mv.Id] = report.Distance
		report.PredecessorMap[mv.Id] = &sv

		sv = mv
	}

	return &report, nil
}

// Dijkstra Traverses the graph using Dijkstra's Algorithm
func Dijkstra(graph Graph, s string, e string) (*AlgorithmReport, error) {
	var sv Vertex
	var ev Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}
	if ev, ok = graph.Vertices[e]; !ok {
		return nil, errors.New(fmt.Sprintf("Ending Vertex: %s is not in Graph.", e))
	}

	report := AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      &ev,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := BlankQueue()
	heap.Push(&queue, &QueueItem{
		Value:    sv,
		Priority: 0,
	})

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*QueueItem)
		cv := cq.Value.(Vertex)

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
				heap.Push(&queue, &QueueItem{
					Value:    *edge.ConnectedVertex,
					Priority: newDist,
				})
			}
		}
	}

	return &report, nil
}

/* Deprecated Implementation */

func _Dijkstra(graph Graph, s string, e string) (*AlgorithmReport, error) {
	var sv Vertex
	var ev Vertex
	var ok bool

	if sv, ok = graph.Vertices[s]; !ok {
		return nil, errors.New(fmt.Sprintf("Starting Vertex: %s is not in Graph.", s))
	}
	if ev, ok = graph.Vertices[e]; !ok {
		return nil, errors.New(fmt.Sprintf("Ending Vertex: %s is not in Graph.", e))
	}

	report := AlgorithmReport{
		StartVertex:    &sv,
		EndVertex:      &ev,
		Distance:       math.Inf(1),
		DistanceMap:    map[string]float64{s: 0},
		PredecessorMap: map[string]*Vertex{s: nil},
		VisitMap:       map[string]bool{s: false},
	}

	queue := BlankQueue()
	heap.Push(&queue, &QueueItem{
		Value:    sv,
		Priority: 0,
	})

	for !queue.IsEmpty() {
		cq := heap.Pop(&queue).(*QueueItem)
		cv := cq.Value.(Vertex)

		if report.VisitMap[cv.Id] {
			continue
		}

		report.VisitMap[cv.Id] = true
		for _, edge := range graph.Edges[cv] {
			newDist := cq.Priority + edge.Weight
			if dist, ok := report.DistanceMap[edge.ConnectedId]; !ok || newDist < dist {
				report.DistanceMap[edge.ConnectedId] = newDist
				report.PredecessorMap[edge.ConnectedId] = &cv
				heap.Push(&queue, &QueueItem{
					Value:    *edge.ConnectedVertex,
					Priority: newDist,
				})
			}
		}
	}

	if report.VisitMap[ev.Id] {

		report.Distance = report.DistanceMap[ev.Id]

		pv := &ev
		for pv != nil {
			report.PredecessorChain = append(report.PredecessorChain, *pv)
			tv := report.PredecessorMap[pv.Id]
			pv = tv
		}
	}

	return &report, nil
}
