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
			if !SliceEq(slice, expected) {
				test.Errorf("expected %+v, got %+v", expected, slice)
			}
		})
	}

	t("plain fill", 3, []int{42, 42, 42})
	t("zero elements", 0, []int{})
}

func TestReverse(test *testing.T) {
	var t = func(name string, slice, expected Slice[int]) {
		test.Run(name, func(test *testing.T) {
			slice.Reverse()
			if !SliceEq(slice, expected) {
				test.Errorf("expected %+v, got %+v", expected, slice)
			}
		})
	}
	t("1 2 3", []int{1, 2, 3}, []int{3, 2, 1})
	t("1 2 3 4", []int{1, 2, 3, 4}, []int{4, 3, 2, 1})
	t("empty", nil, nil)
}
