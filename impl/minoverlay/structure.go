package minoverlay

import "github.com/abhishtagatya/graf"

type OverlayGraph struct {
	Vertices map[string]graf.Vertex
	Edges    map[graf.Vertex][]graf.Edge

	OverlayMap map[string]map[string]OverlayPair
}

func BlankOverlayGraph() OverlayGraph {
	graph := OverlayGraph{
		Vertices:   map[string]graf.Vertex{},
		Edges:      map[graf.Vertex][]graf.Edge{},
		OverlayMap: make(map[string]map[string]OverlayPair),
	}

	return graph
}

func (ogr *OverlayGraph) AddEdge(u, v graf.Vertex, weight float64) graf.Edge {
	uvEdge := graf.Edge{
		Weight:          weight,
		ConnectedId:     v.Id,
		ConnectedVertex: &v,
	}
	ogr.Edges[u] = append(ogr.Edges[u], uvEdge)
	return uvEdge
}

func (ogr *OverlayGraph) AddOverlayEntry(u, v graf.Vertex, ovPair OverlayPair) OverlayPair {
	innerMap := make(map[string]OverlayPair)
	innerMap[v.Id] = ovPair

	ogr.OverlayMap[u.Id] = innerMap

	return ogr.OverlayMap[u.Id][v.Id]
}

func removeElementByValue(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice // If the value isn't found, return the original slice
}
