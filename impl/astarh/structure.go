package astarh

import (
	"container/heap"
	"graf"
)

// HeuristicAlgorithmReport Stores information of algorithm runs
type HeuristicAlgorithmReport struct {
	StartVertex *graf.Vertex `json:"startVertex"`
	EndVertex   *graf.Vertex `json:"endVertex"`

	Distance float64 `json:"minDistance"`

	DistanceMap    map[string]float64      `json:"distanceMap"`
	HeuristicMap   map[string]float64      `json:"heuristicMap"`
	PredecessorMap map[string]*graf.Vertex `json:"predecessorMap"`
	VisitMap       map[string]bool         `json:"visitMap"`

	PredecessorChain []graf.Vertex `json:"predecessorChain"`
}

type HeuristicQueueItem struct {
	Value    interface{}
	Weight   float64
	Priority float64
	Index    int
}

type MinHeuristicQueue []*HeuristicQueueItem

func (mhq MinHeuristicQueue) Len() int {
	return len(mhq)
}

func (mhq MinHeuristicQueue) IsEmpty() bool {
	return len(mhq) == 0
}

func (mhq MinHeuristicQueue) Less(i, j int) bool {
	return mhq[i].Priority < mhq[j].Priority
}

func (mhq MinHeuristicQueue) Swap(i, j int) {
	mhq[i], mhq[j] = mhq[j], mhq[i]
	mhq[i].Index = i
	mhq[j].Index = j
}

func (mhq *MinHeuristicQueue) Push(x interface{}) {
	item := x.(*HeuristicQueueItem)
	item.Index = len(*mhq)
	*mhq = append(*mhq, item)
}

func (mhq *MinHeuristicQueue) Pop() interface{} {
	old := *mhq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*mhq = old[0 : n-1]
	return item
}

func BlankHeuristicQueue() MinHeuristicQueue {
	mhq := make(MinHeuristicQueue, 0)
	heap.Init(&mhq)
	return mhq
}
