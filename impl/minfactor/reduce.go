package minfactor

import (
	"container/heap"
	"github.com/abhishtagatya/graf"
)

// ReduceGraphEdge Reduces the Graph Edge to Minimal Amount of Edges Required
func ReduceGraphEdge(graph graf.Graph) (*graf.Graph, error) {
	redGraph := graf.BlankGraph()

	for sid := range graph.Vertices {
		sv := graph.Vertices[sid]

		queue := graf.BlankMinPriorityQueue()
		heap.Push(&queue, &graf.QueueItem{
			Value:    sv,
			Priority: 0,
		})

		visitMap := map[string]bool{sv.Id: false}
		distanceMap := map[string]float64{sv.Id: 0}

		for !queue.IsEmpty() {
			cq := heap.Pop(&queue).(*graf.QueueItem)
			cv := cq.Value.(graf.Vertex)

			if visitMap[cv.Id] {
				continue
			}

			visitMap[cv.Id] = true
			redGraph.Vertices[cv.Id] = cv

			for _, edge := range graph.Edges[cv] {
				newDist := cq.Priority + edge.Weight
				if dist, ok := distanceMap[edge.ConnectedId]; !ok || newDist < dist {
					distanceMap[edge.ConnectedId] = newDist
					heap.Push(&queue, &graf.QueueItem{
						Value:    *edge.ConnectedVertex,
						Priority: newDist,
					})

					if !graf.ContainsEdge(redGraph.Edges[cv], edge) {
						redGraph.AddEdge(cv, *edge.ConnectedVertex, edge.Weight)
					}
				}
			}
		}
	}

	return &redGraph, nil
}
