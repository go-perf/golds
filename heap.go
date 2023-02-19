package golds

// Heap data structure.
type Heap[T any] struct {
	data []T
	cmp  func(a, b T) int
}

// NewHeap instantiates a new heap with a given size.
func NewHeap[T any](size int, cmp func(a, b T) int) *Heap[T] {
	h := &Heap[T]{
		data: make([]T, 0, size),
		cmp:  cmp,
	}
	return h
}

// NewHeapFrom instantiates a new heap with a given size.
func NewHeapFrom[T any](v []T, cmp func(a, b T) int) *Heap[T] {
	h := &Heap[T]{
		data: make([]T, len(v)),
		cmp:  cmp,
	}
	h.Build(v)
	return h
}

// Size of the heap.
func (h *Heap[T]) Size() int {
	return len(h.data)
}

// Reset the structure. Remove all the elements.
func (h *Heap[T]) Reset() {
	h.data = h.data[:0]
}

// Build pushes all items from values to the heap
func (h *Heap[T]) Build(values []T) {
	size := len(values)
	if len(h.data) < size {
		h.data = make([]T, size)
	}

	copy(h.data, values)
	for i := size/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// Push element to the heap.
func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

// PushMany elements to the heap.
func (h *Heap[T]) PushMany(values ...T) {
	for _, v := range values {
		h.Push(v)
	}
}

// Pop removes and returns top element of the heap
func (h *Heap[T]) Pop() (value T, ok bool) {
	if h.Size() == 0 {
		return zeroOf[T](), false
	}
	size := len(h.data) - 1
	h.swap(0, size)
	h.down(0)
	value, h.data = h.data[size], h.data[:size]
	return value, true
}

// PopMany removes and returns top k elements of the heap
func (h *Heap[T]) PopMany(k int) (values []T, ok bool) {
	if h.Size() == 0 {
		return nil, false
	}

	k = min(k, len(h.data))
	values = make([]T, k)
	for i := 0; i < k; i++ {
		value, _ := h.Pop()
		values[i] = value
	}
	return values, true
}

// Top returns top element of the heap
func (h *Heap[T]) Top() (value T, ok bool) {
	if h.Size() == 0 {
		return zeroOf[T](), false
	}
	return h.data[0], true
}

// Values that are in the heap.
func (h *Heap[T]) Values() []T {
	return h.data[:]
}

// down pushes element down in the heap-tree
func (h *Heap[T]) down(i int) {
	size := len(h.data) - 1
	for {
		j := 2*i + 1
		if j >= size {
			break
		}
		if j2 := j + 1; j2 < size && h.cmp(h.data[j2], h.data[j]) < 0 {
			j = j2
		}
		if h.cmp(h.data[i], h.data[j]) < 0 {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// up pushes element up in the heap-tree
func (h *Heap[T]) up(i int) {
	for {
		j := (i - 1) / 2
		if h.cmp(h.data[i], h.data[j]) >= 0 {
			break
		}
		h.swap(j, i)
		i = j
	}
}

// swap swaps elements on the given indexes
func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func zeroOf[T any]() T {
	var zero T
	return zero
}
