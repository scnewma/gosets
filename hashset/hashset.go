package hashset

import (
	"fmt"
	"strings"
)

// Set represents a hash set. It is implemented with stdlib map with `struct{}`
// as the value.
// All elements in the set must implement comparable.
type Set[T comparable] struct {
	inner map[T]struct{}
}

// New creates a hash set. Passing nil for elems creates an empty hash set.
func New[T comparable](elems []T) *Set[T] {
	s := &Set[T]{
		inner: make(map[T]struct{}, len(elems)),
	}
	for _, i := range elems {
		s.Insert(i)
	}
	return s
}

// Insert adds a new value to the set. Returns true if this is a new value to
// the set, false if it was already in the set.
func (s *Set[T]) Insert(val T) bool {
	if s.Contains(val) {
		return false
	}
	s.inner[val] = struct{}{}
	return true
}

// Merge adds all elements from o to this set.
func (s *Set[T]) Merge(o *Set[T]) {
	for k := range o.inner {
		s.Insert(k)
	}
}

// Remove removes the value from the set. Returns true if the value was present
// in the set, false otherwise.
func (s *Set[T]) Remove(val T) bool {
	present := false
	if s.Contains(val) {
		present = true
	}
	delete(s.inner, val)
	return present
}

// Contains returns true if the set contains the given value.
func (s *Set[T]) Contains(val T) bool {
	_, ok := s.inner[val]
	return ok
}

// Elems returns a list of elements present in the set. They are not returned
// in any particular order.
// nil is returned for an empty set
func (s *Set[T]) Elems() []T {
	if s.Len() == 0 {
		return nil
	}
	elems := make([]T, len(s.inner))
	i := 0
	for k := range s.inner {
		elems[i] = k
		i++
	}
	return elems
}

// Diff returns a new set of the values that are in this hash set, but not
// present in o.
func (s *Set[T]) Diff(o *Set[T]) *Set[T] {
	diff := New[T](nil)
	for k := range s.inner {
		if o.Contains(k) {
			continue
		}
		diff.Insert(k)
	}
	return diff
}

// SymDiff returns a new set of values that are in set or o, but not both.
func (s *Set[T]) SymDiff(o *Set[T]) *Set[T] {
	diff := s.Diff(o)
	diff.Merge(o.Diff(s))
	return diff
}

// Intersect returns a new set representing the intersection of the two sets
// (i.e. the values are in both set and o).
func (s *Set[T]) Intersect(o *Set[T]) *Set[T] {
	intersect := New[T](nil)
	for k := range s.inner {
		if !o.Contains(k) {
			continue
		}
		intersect.Insert(k)
	}
	return intersect
}

// Union returns a new set containing all values in the current set and all
// values in o, without duplicates.
func (s *Set[T]) Union(o *Set[T]) *Set[T] {
	// optimized to add all of the elements from the longest set, then add
	// elements from longest.Diff(shortest). this allows us to know the final
	// size ahead of time and preallocate the inner map.
	s1, s2 := s, o
	if s.Len() < o.Len() {
		s1 = o
		s2 = s
	}
	diff := s2.Diff(s1)

	union := &Set[T]{inner: make(map[T]struct{}, s1.Len()+diff.Len())}
	for k := range s1.inner {
		union.Insert(k)
	}
	for k := range diff.inner {
		union.Insert(k)
	}
	return union
}

// Disjoint returns true if the sets have no elements in common.
func (s *Set[T]) Disjoint(o *Set[T]) bool {
	for k := range s.inner {
		if o.Contains(k) {
			return false
		}
	}
	return true
}

// Subset returns true if this set is a subset of o, meaning o contains all
// elements in this set.
func (s *Set[T]) Subset(o *Set[T]) bool {
	if s.Len() > o.Len() {
		return false
	}

	for k := range s.inner {
		if !o.Contains(k) {
			return false
		}
	}
	return true
}

// Superset returns true if this set is a superset of o, meaning this set
// contains all elements in set o.
func (s *Set[T]) Superset(o *Set[T]) bool {
	return o.Subset(s)
}

// Len returns the length of the set.
func (s *Set[T]) Len() int {
	return len(s.inner)
}

// Clear removes all elements from the hash set, retaining the existing
// capacity.
func (s *Set[T]) Clear() {
	for k := range s.inner {
		delete(s.inner, k)
	}
}

// Equal returns true if both sets contain the exact same elements.
func (s *Set[T]) Equal(o *Set[T]) bool {
	if s.Len() != o.Len() {
		return false
	}
	for k := range s.inner {
		if !o.Contains(k) {
			return false
		}
	}
	return true
}

// String returns a string representation of the set.
func (s *Set[T]) String() string {
	var sb strings.Builder
	sb.WriteByte('{')
	i := 0
	for k := range s.inner {
		sb.WriteString(fmt.Sprintf("%v", k))
		if i < len(s.inner)-1 {
			sb.WriteByte(',')
		}
		i++
	}
	sb.WriteByte('}')
	return sb.String()
}
