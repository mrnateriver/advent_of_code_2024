package shared

import (
	"container/heap"
	"fmt"
)

type priorityEntry[T any] struct {
	value    T
	priority int
}

type PriorityQueue[T any] []priorityEntry[T]

func MakePriorityQueue[T any]() PriorityQueue[T] {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)
	return pq
}

func (pq *PriorityQueue[T]) PushEntry(value T, priority int) {
	heap.Push(pq, priorityEntry[T]{value, priority})
}

func (pq PriorityQueue[T]) PeekEntry() T {
	if len(pq) < 1 {
		panic(fmt.Errorf("attempt to poll empty priority queue"))
	}

	return pq[len(pq)-1].value
}

func (pq *PriorityQueue[T]) PollEntry() T {
	entry := heap.Pop(pq).(priorityEntry[T]) // Pop in heap is poll in priority queue, because lower int value - the higher the priority, like #1
	return entry.value
}

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(priorityEntry[T])
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
