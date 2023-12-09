package graf

// AlgorithmReport Stores information of algorithm runs for Dijkstra
type AlgorithmReport struct {
	StartVertex *Vertex `json:"startVertex"`
	EndVertex   *Vertex `json:"endVertex"`

	Distance float64 `json:"minDistance"`

	DistanceMap    map[string]float64 `json:"distanceMap"`
	PredecessorMap map[string]*Vertex `json:"predecessorMap"`
	VisitMap       map[string]bool    `json:"visitMap"`

	PredecessorChain []Vertex `json:"predecessorChain"`
}
