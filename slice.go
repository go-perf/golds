package golds

// Slice is a sequence of values.
// Basically it's an utility wrapper around a plain slice.
type slice[E any] []E

// NewSlice puts provided values into resulting slice.
//export
func newSlice[E any] (vv ... E) slice[E] {
	return slice[E](vv)
}

// Repeat returns a new Slice[E] with n copies of v.
func repeat[E any](n int, v E) slice[E] {
	switch{
	case n == 0:
		return nil
	case n < 0:
		panic("golds.Repeat: negative Repeat count")
	}
	var s = make(slice[E], n)
	// TODO: use copy to fill slice for better performance
	for i := range s {
		s[i] = v
	}
	return s
} 

// SliceIterFn returns a new Slice[E] with results of fn(i) for i := range [0, n].
func sliceIterFn[E any](n int, fn func(int) E) slice[E] {
	switch{
	case n == 0:
		return nil
	case n < 0:
		panic("golds.Repeat: negative Repeat count")
	}
	var s = make(slice[E], n)
	for i := range s {
		s[i] = fn(i)
	}
	return s
}

// Len returns the number of elements in the slice.
func(s slice[E]) Len() int { return len(s) }

// Cap returns capacity of underlying slice.
func(s slice[E]) Cap() int { return cap(s) }

// Index returns element by the index.
// If index is negative, then returns element len-index.
// Panics if index is out of [-len, len] range.
func(s slice[E]) Index(i int) E {
	if i < 0 {
		return s[s.Len() - i]
	}
	return s[i]
} 

// Swap two elements in the slice.
// If i or j are negative, then uses elements len-i(j).
func(s slice[E]) Swap(i, j int) { 
	if i < 0 {
		i = s.Len() - i
	}
	if j < 0 {
		j = s.Len() - j
	}
	s[i], s[j] = s[j], s[i]
}

// Count returns the number of elements e, where fn(e) is true.
func(s slice[E]) Count(fn func(v E) bool) int {
	var i int
	for _, v := range s {
		if fn(v) {
			i++
		}
	}
	return i
} 

// Filter returns slice of elements e, where fn(e) is true.
// fn can be called multiple times for each element.
func(s slice[E]) Filter(fn func(v E) bool) slice[E] {
	var filtered = make(slice[E], 0, s.Count(fn))
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func(s slice[E]) FilterInPlace(fn func(v E) bool) {
	var i int
	for _, v := range s {
		if fn(v) {
			s[i] = v
			i++
		}
	}
	var empty E
	s[i:].Fill(empty)
}

// Select returns gets elements by indexes and puts them into a new slice.
// If index is negative, then len-index element will be used.
// Example:
//	Slice[int]{1, 2, 3}.Select(-1, 0, 2) -> Slice[int]{3, 1, 2}
func(s slice[E]) Select(indexes ... int) slice[E] {
	var selected = make(slice[E], 0, len(indexes))
	for _, i := range indexes {
		selected = append(selected, s.Index(i))
	}
	return selected
}

// Apply creates a new slice with mapped values.
func(s slice[E]) Apply(fn func(v E) E) slice[E] {
	var result = make(slice[E], 0, s.Len())
	for _, v := range s {
		result = append(result, fn(v))
	}
	return s
} 

// Append new elements to the slice in place.
func(s *slice[E]) Append(vv ...E) {
	*s = append(*s, vv...)
}


// Pop returns returns the last element and removes it from the slice.
// If the slice is empty, then returns false.
func(s *slice[E]) Pop() (E, bool) {
	var empty E
	var n = s.Len()
	if n == 0 {
		return empty, false
	}
	var v = (*s)[n-1]
	(*s)[n-1] = empty
	*s = (*s)[:n-1]
	return v, true
}

func(s slice[E]) Copy() slice[E] {
	var cp = make(slice[E], s.Len())
	copy(cp, s)
	return cp
}

func(s slice[E]) CopyWith(fn func(E) E) slice[E] {
	var cp = make(slice[E], s.Len())
	for i, v := range s {
		cp[i] = fn(v)
	}
	return cp
}

// Fill slice using provided value.
func(s slice[E]) Fill(v E) {
	var n = s.Len()
	if n == 0 {
		return
	}
	s[0] = v
	for i := 1; i < n; i *= 2 {
		copy(s[:], s[:i])
	}
}

// FillWith uses results of fn to fill the slice.
func(s slice[E]) FillWith(fn func() E) {
	for i := range s {
		s[i] = fn()
	}
}

// Insert value at i-position, shifting elements to the end of slice.
// Panics if the index is out of range.
// 	Slice{1, 2, 3}.Insert(1, 100) -> Slice{1, 100, 2, 3}
func(s *slice[E]) Insert(i int, v E) {
	var sl = *s
	*s = append(sl[:i+1], sl[i:]...)
	(*s)[i] = v
}

// Delete value at i-position, shifting elements to the begining of slice.
// Panics if the index is out of range.
// 	Slice{1, 2, 3}.Delete(1) -> Slice{2, 3}
func(s *slice[E]) Delete(i int) {
	var sl = *s
	sl = append(sl[:i], sl[i+1:]...)

	var empty E
	(*s)[s.Len()-1] = empty
	*s = sl
}

// DeleteNoOrder exchanges the i-th and the last element 
// of the slice and cuts the last, now duplicated, element.
func(s *slice[E]) DeleteNoOrder(i int) {
	var n = s.Len()
	var sl = *s
	sl[i] = sl[n-1]

	var empty E
	sl[n-1] = empty
	*s = sl[:n-1]
}
