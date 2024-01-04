package graf

import "math"

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

type GraphMatrix map[string]map[string]float64

// AddVertex Adding a Vertex to the Graph
func (gr *Graph) AddVertex(v string) Vertex {
	var vf Vertex
	var ok bool

	if vf, ok = gr.Vertices[v]; !ok {
		gr.Vertices[v] = Vertex{Id: v}
		vf = gr.Vertices[v]
	}

	return vf
}

// AddEdge Adding an Edge to the Graph
func (gr *Graph) AddEdge(u, v Vertex, weight float64) Edge {
	uvEdge := Edge{
		Weight:          weight,
		ConnectedId:     v.Id,
		ConnectedVertex: &v,
	}
	gr.Edges[u] = append(gr.Edges[u], uvEdge)
	return uvEdge
}

// ToAdjMatrix Transform Adjacency-List Representation (Default) to an Adjacency-Matrix Representation
func ToAdjMatrix(graph *Graph) GraphMatrix {
	matrix := make(GraphMatrix)

	for vi, _ := range graph.Vertices {
		innerMap := make(map[string]float64)
		for vk, _ := range graph.Vertices {
			innerMap[vk] = math.Inf(1)
			if vi == vk {
				innerMap[vk] = 0
			}
		}
		matrix[vi] = innerMap
	}

	for vi, vObj := range graph.Vertices {
		for _, edge := range graph.Edges[vObj] {
			matrix[vi][edge.ConnectedId] = edge.Weight
		}
	}

	return matrix
}

// CountEdge Count the Edge of a Graph
func CountEdge(graph *Graph) int {
	ec := 0
	for _, edge := range graph.Edges {
		ec += len(edge)
	}

	return ec
}

// ContainsVertex Checks if a Graph contains a searched Vertex
func ContainsVertex(list []string, v string) bool {
	for _, i := range list {
		if v == i {
			return true
		}
	}

	return false
}

// ContainsEdge Checks if a Graph contains a searched Edge
func ContainsEdge(edges []Edge, edge Edge) bool {
	for _, e := range edges {
		if e.ConnectedId == edge.ConnectedId && e.ConnectedVertex.Id == edge.ConnectedVertex.Id && e.Weight == edge.Weight {
			return true
		}
	}

	return false
}

// BlankGraph Create a Blank Graph
func BlankGraph() Graph {
	graph := Graph{
		Vertices: map[string]Vertex{},
		Edges:    map[Vertex][]Edge{},
	}

	return graph
}
