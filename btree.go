package golds

// no floats, since they are just partially ordered, thanks to NaN
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~string
}

type BTree[K Ordered, V any] struct {
	// TODO: add arena node and kv allocator here
	root *btreeNode[K, V]
}

func(bt *BTree[K, V]) Find(key K) (V, bool) {

	var empty V
	return empty, false
}

type orderedKV[K Ordered, V any] struct {
	key   K
	value V
}

type btreeNode[K Ordered, V any] struct {
	kv       []orderedKV[K, V]
	children []*btreeNode[K, V]
}
