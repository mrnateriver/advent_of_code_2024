package shared

import "cmp"

type BST[T cmp.Ordered] struct {
	Left  *BST[T]
	Value T
	Right *BST[T]
}

func (t *BST[T]) walk(ch chan T) {
	if t == nil {
		return
	}

	t.Left.walk(ch)
	ch <- t.Value
	t.Right.walk(ch)
}

func (t *BST[T]) WalkerDfs() <-chan T {
	ch := make(chan T)

	go func() {
		t.walk(ch)
		close(ch)
	}()

	return ch
}

func (t *BST[T]) Insert(v T) *BST[T] {
	if t == nil {
		return &BST[T]{nil, v, nil}
	}

	if v < t.Value {
		t.Left = t.Left.Insert(v)
		return t
	}

	t.Right = t.Right.Insert(v)
	return t
}

func (t *BST[T]) Size() int {
	if t == nil {
		return 0
	}

	if t.Left == nil && t.Right == nil {
		return 1
	}

	size := 1
	if t.Left != nil {
		size += t.Left.Size()
	}
	if t.Right != nil {
		size += t.Right.Size()
	}

	return size
}
