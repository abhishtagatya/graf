package graf

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

func ContainsEdge(edges []Edge, edge Edge) bool {
	for _, e := range edges {
		if e.ConnectedId == edge.ConnectedId && e.ConnectedVertex.Id == edge.ConnectedVertex.Id && e.Weight == edge.Weight {
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
