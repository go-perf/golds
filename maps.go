
// Package maps defines various functions useful with maps of any type.
// Note: Will be replaces by stdlib "package maps" in the future
package golds

// Keys returns the keys of the map m.
// The keys will be an indeterminate order.
func Keys[K comparable, V any](m map[K]V) []K{
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys,k)
	}
	return keys
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[K comparable, V any](m map[K]V) []V{
	values := make([]V, 0, len(m))
	for _,v := range m {
		values = append(values,v)
	}
	return values
}

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[K, V comparable](m1, m2 map[K]V) bool{
	if len(m1) != len(m2) {
		return false
	}
	for e := range m1 {
		if !m2.contains(e) {
			return false
		}
	}
	return true
}

// EqualFunc is like Equal, but compares values using cmp.
// Keys are still compared with ==.
func EqualFunc[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, cmp func(V1, V2) bool) bool{
	panic("unimplemented")
}

// Clear removes all entries from m, leaving it empty.
func Clear[K comparable, V any](m map[K]V){
	m = make(map[K]V, 0)
}

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[K comparable, V any](m map[K]V) map[K]V{
	panic("unimplemented")
}

// Add adds all key/value pairs in src to dst. When a key in src
// is already present in dst, the value in dst will be overwritten
// by the value associated with the key in src.
func Add[K comparable, V any](dst, src map[K]V){
	panic("unimplemented")
}

// Filter deletes any key/value pairs from m for which keep returns false.
func Filter[K comparable, V any](m map[K]V, keep func(K, V) bool){
	panic("unimplemented")
}
