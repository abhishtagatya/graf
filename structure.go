package graf

import (
	"container/heap"
)

/* Graph Structure */

type Edge struct {
	Weight          float64 `json:"weight"`
	ConnectedId     string  `json:"connectedId"`
	ConnectedVertex *Vertex `json:"connectedVertex"`
}

type Vertex struct {
	Id string `json:"id"`
}

// Graph G = (V, E)
type Graph struct {
	Vertices map[string]Vertex `json:"vertices"`
	Edges    map[Vertex][]Edge `json:"edges"`
}

func (gr *Graph) AddVertex(v string) Vertex {
	var vf Vertex
	var ok bool

	if vf, ok = gr.Vertices[v]; !ok {
		gr.Vertices[v] = Vertex{Id: v}
		vf = gr.Vertices[v]
	}

	return vf
}

func (gr *Graph) AddEdge(u, v Vertex, weight float64) Edge {
	uvEdge := Edge{
		Weight:          weight,
		ConnectedId:     v.Id,
		ConnectedVertex: &v,
	}
	gr.Edges[u] = append(gr.Edges[u], uvEdge)
	return uvEdge
}

func CountEdge(graph *Graph) int {
	ec := 0
	for _, edge := range graph.Edges {
		ec += len(edge)
	}

	return ec
}

func ContainsVertex(list []string, v string) bool {
	for _, i := range list {
		if v == i {
			return true
		}
	}

	return false
}

func BlankGraph() Graph {
	graph := Graph{
		Vertices: map[string]Vertex{},
		Edges:    map[Vertex][]Edge{},
	}

	return graph
}

/* Queue Structure */

type QueueItem struct {
	Value    interface{}
	Priority float64
	Index    int
}

type MinPriorityQueue []*QueueItem

func (mpq MinPriorityQueue) Len() int {
	return len(mpq)
}

func (mpq MinPriorityQueue) IsEmpty() bool {
	return len(mpq) == 0
}

func (mpq MinPriorityQueue) Less(i, j int) bool {
	return mpq[i].Priority < mpq[j].Priority
}

func (mpq MinPriorityQueue) Swap(i, j int) {
	mpq[i], mpq[j] = mpq[j], mpq[i]
	mpq[i].Index = i
	mpq[j].Index = j
}

func (mpq *MinPriorityQueue) Push(x interface{}) {
	item := x.(*QueueItem)
	item.Index = len(*mpq)
	*mpq = append(*mpq, item)
}

func (mpq *MinPriorityQueue) Pop() interface{} {
	old := *mpq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*mpq = old[0 : n-1]
	return item
}

func BlankQueue() MinPriorityQueue {
	mpq := make(MinPriorityQueue, 0)
	heap.Init(&mpq)
	return mpq
}
