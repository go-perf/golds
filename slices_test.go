package golds

import (
	"testing"
)

func TestFill(test *testing.T) {
	const x = 42
	var t = func(name string, n int, expected []int) {
		test.Run(name, func(test *testing.T) {
			var slice = make(Slice[int], n)
			slice.Fill(x)
			SliceEq(slice, expected)
		})
	}

	t("plain fill", 3, []int{42, 42, 42})
	t("zero elements", 0, []int{})
}