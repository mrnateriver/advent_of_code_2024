package shared

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) walk(ch chan int) {
	if t == nil {
		return
	}

	t.Left.walk(ch)
	ch <- t.Value
	t.Right.walk(ch)
}

func (t *Tree) Walker() <-chan int {
	ch := make(chan int)

	go func() {
		t.walk(ch)
		close(ch)
	}()

	return ch
}

func (t *Tree) Insert(v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v < t.Value {
		t.Left = t.Left.Insert(v)
		return t
	}

	t.Right = t.Right.Insert(v)
	return t
}

func (t *Tree) Size() int {
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
