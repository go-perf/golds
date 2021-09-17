package golds

import (
	"testing"
)

func TestSliceIndex(test *testing.T) {
	var t = func(name string, slice Slice[string], i int, expected string) {
		test.Run(name, func(test *testing.T) {
			var v = slice.Index(i)
			if v != expected {
				test.Errorf("expected value %q, got %q", v, expected)
			}
		})
	}

	var slice = Slice[string]{"one", "horse", "apple"}
	t("middle element", slice, 1, "horse")
	t("first from the end", slice, -1, "apple")

	test.Run("out of range", testSliceIndexPanics)
}

func testSliceIndexPanics(test *testing.T) {
	var t = func(name string, slice Slice[string], i int) {
		test.Run(name, func(test *testing.T) {
			assertPanic(test, func() {
				_ = slice.Index(i)
			})
		})
	}

	var slice = Slice[string]{"one", "horse", "apple"}
	t("middle element", slice, 3)
	t("third from the end", slice, -4)
}

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

func assertPanic(test *testing.T, fn func()) {
	test.Helper()
	func() {
		defer requireRecover(test)
		fn()
	}()
}

func requireRecover(test *testing.T) {
	var v = recover()
	if v == nil {
		test.Errorf("a panic is expected")
		return
	}
	test.Logf("recovered: %v", v)
}
