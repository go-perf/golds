package golds

func zeroOf[T any]() T {
	var zero T
	return zero
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
