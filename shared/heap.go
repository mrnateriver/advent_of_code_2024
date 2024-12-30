package shared

import (
	"container/heap"
	"fmt"
)

type Heap[T any] struct {
	values      []T
	compareLess func(a, b T) bool
}

func MakeHeap[T any](less func(a, b T) bool) Heap[T] {
	pq := Heap[T]{[]T{}, less}
	heap.Init(&pq)
	return pq
}

func (pq *Heap[T]) PushEntry(value T) {
	heap.Push(pq, value)
}

func (pq Heap[T]) PeekEntry() T {
	l := len(pq.values)
	if l < 1 {
		panic(fmt.Errorf("attempt to peek empty heap"))
	}

	return pq.values[l-1]
}

func (pq *Heap[T]) PopEntry() T {
	return heap.Pop(pq).(T)
}

func (pq Heap[T]) Len() int { return len(pq.values) }

func (pq Heap[T]) Less(i, j int) bool {
	return pq.compareLess(pq.values[i], pq.values[j])
}

func (pq Heap[T]) Swap(i, j int) {
	pq.values[i], pq.values[j] = pq.values[j], pq.values[i]
}

func (pq *Heap[T]) Push(x any) {
	pq.values = append(pq.values, x.(T))
}

func (pq *Heap[T]) Pop() any {
	old := pq.values
	n := len(old)
	item := old[n-1]
	pq.values = old[0 : n-1]
	return item
}
