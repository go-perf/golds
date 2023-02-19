package golds

import (
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap(0, intCmp)

	if h == nil {
		t.Error("cannot instantiate Heap")
	}

	_, ok1 := h.Top()
	_, ok2 := h.Pop()
	if ok1 || ok2 {
		t.Errorf("expected to be nil")
	}
	if values, ok := h.PopMany(100); ok || values != nil {
		t.Errorf("expected to be nil")
	}

	h.Push(100)
	for i := 0; i < 10; i++ {
		h.Push(i)
	}

	if value := h.Size(); value != 11 {
		t.Errorf("want size 11, got %v", value)
	}

	tmp := h.Values()
	values := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		values[i] = tmp[i]
	}
	sort.Sort(sort.IntSlice(values))

	for i := 0; i < 10; i++ {
		if values[i] != i {
			t.Errorf("want %v, got %v", i, values[i])
		}
	}

	for i := 0; i < 10; i++ {
		value, ok := h.Pop()
		if !ok || value != i {
			t.Errorf("incorrect value, expected %v got %v", i, value)
		}
	}

	if value, ok := h.Top(); !ok || value != 100 {
		t.Errorf("expected 100, got %v", value)
	}

	h.PushMany(10, 20, 30)
	for i := 1; i <= 3; i++ {
		value, ok := h.Pop()
		if !ok || value != i*10 {
			t.Errorf("incorrect value, expected %v got %v", i*10, value)
		}
	}

	h.PushMany(10, 20)
	values2, ok := h.PopMany(4)
	if !ok || len(values2) != 3 {
		t.Errorf("want size %v, got %v", 3, len(values2))
	}
}

func TestHeapBuild(t *testing.T) {
	h := NewHeap(10, intCmp)
	if h == nil {
		t.Error("cannot instantiate Heap")
	}

	values := make([]int, 100)
	for i := 0; i < 100; i++ {
		if i <= 50 {
			values[i] = i
		} else {
			values[i] = 100 - i
		}
	}
	h.Build(values)

	if value := h.Size(); value != 100 {
		t.Errorf("expected size %v, got %v", 100, value)
	}

	h.Pop()

	for i := 1; i < 50; i++ {
		value1, ok1 := h.Pop()
		value2, ok2 := h.Pop()
		if !ok1 || !ok2 || value1 != i || value2 != i {
			t.Errorf("incorrect values, expected %v, got %v and %v", i, value1, value2)
		}
	}

	h.Reset()
	if h.Size() != 0 {
		t.Errorf("should be empty")
	}
}

func BenchmarkHeapPush(b *testing.B) {
	h := NewHeap(1000, intCmp)
	if h == nil {
		b.Error("cannot instantiate Heap")
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(i)
	}

	if h.Size() != b.N {
		b.Fatalf("got %d want %d", h.Size(), b.N)
	}
}

func BenchmarkHeapPushPop(b *testing.B) {
	h := NewHeap(1000, intCmp)
	if h == nil {
		b.Error("cannot instantiate Heap")
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Push(i)
		if _, ok := h.Pop(); !ok {
			b.Fatal()
		}
	}

	if h.Size() != 0 {
		b.Fatalf("got %d want %d", h.Size(), 0)
	}
}

func intCmp(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
