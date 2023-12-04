package graf

import (
	"container/heap"
)

/* Fully Auxiliary Examples to Specific Problem Implementations. */

type AuxTuple struct {
	U string
	V string
}

func ContainsVertex(list []string, v string) bool {
	for _, i := range list {
		if v == i {
			return true
		}
	}

	return false
}

func ComputeContainers(graph *Graph) map[AuxTuple][]string {
	aMap := make(map[string]AuxTuple)
	auxContainer := make(map[AuxTuple][]string)

	for sid := range graph.Vertices {
		sv := graph.Vertices[sid]

		distanceMap := map[string]float64{sid: 0}

		queue := BlankQueue()
		heap.Push(&queue, &QueueItem{
			Value:    sv,
			Priority: 0,
		})

		for !queue.IsEmpty() {
			cq := heap.Pop(&queue).(*QueueItem)
			cv := cq.Value.(Vertex)

			if cv != sv {
				if !ContainsVertex(auxContainer[aMap[cv.Id]], cv.Id) {
					auxContainer[aMap[cv.Id]] = append(auxContainer[aMap[cv.Id]], cv.Id)
				}
			}

			for _, edge := range graph.Edges[cv] {
				newDist := cq.Priority + edge.Weight
				if dist, ok := distanceMap[edge.ConnectedId]; !ok || newDist < dist {
					distanceMap[edge.ConnectedId] = newDist
					heap.Push(&queue, &QueueItem{
						Value:    *edge.ConnectedVertex,
						Priority: newDist,
					})

					if cv == sv {
						aMap[edge.ConnectedId] = AuxTuple{U: sv.Id, V: edge.ConnectedId}
					} else {
						aMap[edge.ConnectedId] = aMap[cv.Id]
					}
				}
			}

		}
	}

	return auxContainer
}
