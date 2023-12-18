package floydwarshall

import (
	"fmt"
	"graf"
	"math"
)

// FloydWarshall Implementation of the Floyd-Warshall Algorithm for the All-Pairs Problem
func FloydWarshall(graph graf.Graph) (graf.GraphMatrix, error) {
	dMat := graf.ToAdjMatrix(&graph)
	fmt.Println(dMat)

	count := 0
	for k, _ := range graph.Vertices {
		fmt.Println(k, count, len(graph.Vertices))
		for i, _ := range graph.Vertices {
			for j, _ := range graph.Vertices {
				if dMat[i][k] != math.Inf(1) && dMat[k][j] != math.Inf(1) && (dMat[i][k]+dMat[k][j]) < dMat[i][j] {
					dMat[i][j] = dMat[i][k] + dMat[k][j]
				}
			}
		}
		count += 1
	}

	return dMat, nil
}
