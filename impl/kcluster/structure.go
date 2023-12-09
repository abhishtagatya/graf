package kcluster

var XHCContentCluster = "s %s C%d\n"
var XHCContentEdge = "q C%d C%d %s %s %f\n"

var XHCBoilerplate = []string{
	"c Hierarchical Cluster (Aux Graph)\n",
	"c Made with Graf (Graph Algorithms Library in Go)\n",
	"c https://github.com/abhishtagatya/graf\n",
}

type EdgeCluster struct {
	Weight           float64        `json:"weight"`
	ConnectedId      string         `json:"connectedId"`
	ConnectedCluster *VertexCluster `json:"connectedCluster"`

	VertexA string `json:"vertexA"`
	VertexB string `json:"vertexB"`
}

type VertexCluster struct {
	ClusterId string `json:"clusterId"`
}

type VertexClusterAnnot struct {
	VCluster *VertexCluster
	VertexA  string
	VertexB  string
	Weight   float64
}

type ClusterGraph struct {
	VertexCluster map[string]VertexCluster        `json:"vertexCluster"`
	Edges         map[VertexCluster][]EdgeCluster `json:"edges"`
}

func (cgr *ClusterGraph) AddVertexCluster(v, c string) VertexCluster {
	var vc VertexCluster
	var ok bool

	if vc, ok = cgr.VertexCluster[v]; !ok {
		cgr.VertexCluster[v] = VertexCluster{ClusterId: c}
		vc = cgr.VertexCluster[v]
	}

	return vc
}

func (cgr *ClusterGraph) AddEdge(u, v VertexCluster, sv, ev string, weight float64) EdgeCluster {
	uvEdge := EdgeCluster{
		Weight:           weight,
		ConnectedId:      v.ClusterId,
		ConnectedCluster: &v,
		VertexA:          sv,
		VertexB:          ev,
	}
	cgr.Edges[u] = append(cgr.Edges[u], uvEdge)
	return uvEdge
}

func BlankClusterGraph() ClusterGraph {
	cGraph := ClusterGraph{
		VertexCluster: map[string]VertexCluster{},
		Edges:         map[VertexCluster][]EdgeCluster{},
	}

	return cGraph
}
