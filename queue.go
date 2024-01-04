package graf

import "container/heap"

type QueueItem struct {
	Value    interface{}
	Priority float64
	Index    int
}

type Queue []*QueueItem

func (q *Queue) Push(x interface{}) {
	item := x.(*QueueItem)
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	if len(*q) == 0 {
		return nil
	}

	old := *q
	item := old[0]
	*q = old[1:]
	return item
}

func (q *Queue) Peek() interface{} {
	if len(*q) == 0 {
		return nil
	}

	item := *q
	return item[0]
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func (q Queue) Len() int {
	return len(q)
}

func BlankQueue() Queue {
	q := make(Queue, 0)
	return q
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

func BlankMinPriorityQueue() MinPriorityQueue {
	mpq := make(MinPriorityQueue, 0)
	heap.Init(&mpq)
	return mpq
}
