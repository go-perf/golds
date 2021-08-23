package test

import "testing"

// Testing exponential copy fill slice algo.
func TestFill(test *testing.T) {
	fill := func(s []string, v string) {
		var n = len(s)
		if n == 0 {
			return
		}
		s[0] = v
		for i := 1; i < n; i *= 2 {
			copy(s[i:], s[:i])
		}
	}
	var t = func(name string, n int, v string) {
		var slice = make([]string, n)
		fill(slice, v)
		var expect = repeat(n, v)
		if !equalStrings(slice, expect) {
			test.Errorf("expecting %+v, got %+v", expect, slice)
		}
		test.Log(slice)
	}

	t("happy path", 11, "a")
	t("empty", 0, "a")
}

var bCopy = make([]int, 512)

func BenchmarkCopy(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		var n = len(bCopy)
		if n == 0 {
			return
		}
		bCopy[0] = 42
		for index := 1; index < n; index *= 2 {
			copy(bCopy[index:], bCopy[:index])
		}
	}
}

var bSet = make([]int, 512)

func BenchmarkSet(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		for index := range bSet {
			bSet[index] = 42
		}
	}
}

var bAppend = make([]int, 0, 512)

func BenchmarkAppend(bench *testing.B) {
	for i := 0; i < bench.N; i++ {
		bAppend = bAppend[:0]
		for i := 0; i < cap(bAppend); i++ {
			bAppend = append(bAppend, 42)
		}
	}
}


func repeat(n int, v string) []string {
	var slice = make([]string, 0, n)
	for i := 0; i < n; i++ {
		slice = append(slice, v)
	}
	return slice
}

func equalStrings(a,b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}
