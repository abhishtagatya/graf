package graf

import "container/heap"

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
