package golds

// Heap data structure.
type Heap[T any] struct {
	data []T
	cmp  func(a, b T) bool
}

// NewHeap instantiates a new heap with a given size.
func NewHeap[T any](size int, cmp func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{
		data: make([]T, 0, size),
		cmp:  cmp,
	}
	return h
}

// NewHeapFrom instantiates a new heap with a given size.
func NewHeapFrom[T any](v []T, cmp func(a, b T) bool) *Heap[T] {
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
		h.siftDown(i)
	}
}

// Push element to the heap.
func (h *Heap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.siftUpLast()
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
	h.data[0], h.data[size] = h.data[size], h.data[0]
	h.siftDown(0)
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
	vals := make([]T, len(h.data))
	copy(vals, h.data)
	return vals
}

// down pushes element down in the heap-tree
func (h *Heap[T]) siftDown(i int) {
	data := h.data
	for i >= 0 && i < len(data) {
		j := 2*i + 1
		if j <= 0 || j >= len(data)-1 {
			break
		}
		if j2 := j + 1; j2 < len(data)-1 && j2 >= 0 && h.cmp(data[j2], data[j]) {
			j = j2
		}
		if h.cmp(data[i], data[j]) {
			break
		}
		data[i], data[j] = data[j], data[i]
		i = j
	}
}

// siftUpLast pushes element up in the heap-tree.
func (h *Heap[T]) siftUpLast() {
	data := h.data
	i := uint(len(data)) - 1
	j := (uint(len(data)) - 2) >> 1
	for i > j && i < uint(len(data)) {
		if !h.cmp(data[i], data[j]) {
			break
		}
		data[i], data[j] = data[j], data[i]
		i, j = j, (j-1)>>1
	}
}
