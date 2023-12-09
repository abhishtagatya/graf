package kcluster

import (
	"graf"
	"math"
)

// ClusterTuple Tuple Type for Cluster Vertices
type ClusterTuple struct {
	U int
	V int
}

// ClusterEdge Edge of Clusters for Computation
type ClusterEdge struct {
	UVertex string
	VVertex string
	Weight  float64
}

// ComputeKCluster Compute K-Cluster for a given Graph
func ComputeKCluster(graph *graf.Graph, k int) [][]string {
	clusters := make([][]string, len(graph.Vertices))

	count := 0
	for v := range graph.Vertices {
		clusters[count] = append(clusters[count], v)
		count += 1
	}

	for len(clusters) > k {
		minWeight := math.Inf(1)
		var closestCluster []int

		for i := 0; i < len(clusters); i++ {
			for j := i + 1; j < len(clusters); j++ {
				interCluster := ComputeClusterWeight(graph, clusters[i], clusters[j])
				if interCluster.Weight < minWeight {
					minWeight = interCluster.Weight
					closestCluster = []int{i, j}
				}
			}
		}

		mergedCluster := append(clusters[closestCluster[0]], clusters[closestCluster[1]]...)
		clusters = append(clusters[:closestCluster[0]], clusters[closestCluster[0]+1:]...)
		clusters = append(clusters[:closestCluster[1]-1], clusters[closestCluster[1]:]...)
		clusters = append(clusters, mergedCluster)

	}

	return clusters

}

// ComputeClusterWeight Compute the Min-Weight between Two Clusters
func ComputeClusterWeight(graph *graf.Graph, clusterA []string, clusterB []string) ClusterEdge {
	uVertex := ""
	vVertex := ""
	minWeight := math.Inf(1)

	for _, nodeA := range clusterA {
		for _, edge := range graph.Edges[graph.Vertices[nodeA]] {
			if graf.ContainsVertex(clusterB, edge.ConnectedId) && edge.Weight < minWeight {
				minWeight = edge.Weight
				uVertex = nodeA
				vVertex = edge.ConnectedId
			}
		}
	}
	return ClusterEdge{UVertex: uVertex, VVertex: vVertex, Weight: minWeight}
}
