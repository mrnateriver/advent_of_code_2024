package shared

type priorityEntry[T any] struct {
	value    T
	priority int
}

type PriorityQueue[T any] struct {
	heap Heap[priorityEntry[T]]
}

func MakePriorityQueue[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{
		heap: MakeHeap[priorityEntry[T]](func(a, b priorityEntry[T]) bool {
			return a.priority < b.priority
		}),
	}
}

func (pq *PriorityQueue[T]) PushEntry(value T, priority int) {
	pq.heap.Push(priorityEntry[T]{value, priority})
}

func (pq PriorityQueue[T]) PeekEntry() T {
	entry := pq.heap.PeekEntry()
	return entry.value
}

func (pq *PriorityQueue[T]) PollEntry() T {
	// Pop in heap is poll in this priority queue, because lower int value - the higher the priority, like #1
	// This is different in some implementations, but at least this naming disparity makes this implementation
	// consistent with the heap package
	entry := pq.heap.PopEntry()
	return entry.value
}

func (pq PriorityQueue[T]) Len() int { return pq.heap.Len() }
