package hashset

import "testing"

func TestSetBasics(t *testing.T) {
	s := New[int](nil)
	assertEq(t, s.Len(), 0)

	// basic insert / remove
	assert(t, !s.Contains(1))
	assert(t, s.Insert(1))
	assertEq(t, s.Len(), 1)
	assert(t, s.Contains(1))
	assert(t, !s.Insert(1)) // false for duplicates
	assert(t, s.Remove(1))
	assert(t, !s.Remove(1)) // false if not in set

	for i := 0; i < 100; i++ {
		assert(t, s.Insert(i))
	}
	assertEq(t, s.Len(), 100)
	for i := 0; i < 100; i++ {
		assert(t, s.Contains(i))
	}
	s.Clear()
	assertEq(t, s.Len(), 0)
}

func TestSetElems(t *testing.T) {
	s := New[int](nil)
	assertEq(t, len(s.Elems()), 0)

	s.Merge(New([]int{1, 2, 3}))
	elems := s.Elems()
	assertEq(t, len(elems), 3)
	assert(t, sliceContains(elems, 1))
	assert(t, sliceContains(elems, 2))
	assert(t, sliceContains(elems, 3))
}

func TestSetMerge(t *testing.T) {
	t.Run("two empty sets", func(t *testing.T) {
		s := New[int](nil)
		s.Merge(New[int](nil))
		assertEq(t, s.Len(), 0)
	})

	t.Run("into empty set", func(t *testing.T) {
		s := New[int](nil)
		s.Merge(New([]int{1}))
		assertEq(t, s.Len(), 1)
		assert(t, s.Contains(1))
	})

	t.Run("empty set into not empty", func(t *testing.T) {
		s := New([]int{1})
		s.Merge(New[int](nil))
		assertEq(t, s.Len(), 1)
		assert(t, s.Contains(1))
	})

	t.Run("non-overlapping sets", func(t *testing.T) {
		s1 := New([]int{1, 2, 3})
		s2 := New([]int{4, 5, 6})
		s1.Merge(s2)
		assertEq(t, s1.Len(), 6)
		for i := 1; i <= 6; i++ {
			assert(t, s1.Contains(i))
		}
		// s2 should not be modified
		assertEq(t, s2.Len(), 3)
		for i := 4; i <= 6; i++ {
			assert(t, s2.Contains(i))
		}
	})

	t.Run("completely overlapping sets", func(t *testing.T) {
		s1 := New([]int{1, 2, 3})
		s2 := New([]int{1, 2, 3})
		s1.Merge(s2)
		assertEq(t, s1.Len(), 3)
		assertEq(t, s2.Len(), 3)
		for i := 1; i <= 3; i++ {
			assert(t, s1.Contains(i))
			assert(t, s2.Contains(i))
		}
	})

	t.Run("partially overlapping sets", func(t *testing.T) {
		s1 := New([]int{1, 2, 3, 4})
		s2 := New([]int{3, 4, 5, 6})
		s1.Merge(s2)
		assertEq(t, s1.Len(), 6)
		for i := 1; i <= 6; i++ {
			assert(t, s1.Contains(i))
		}
		assertEq(t, s2.Len(), 4)
		for i := 3; i <= 6; i++ {
			assert(t, s2.Contains(i))
		}
	})
}

func TestSetUnion(t *testing.T) {
	t.Run("two empty sets", func(t *testing.T) {
		s1 := New[int](nil)
		s2 := New[int](nil)
		s3 := s1.Union(s2)
		assertEq(t, s1.Len(), 0)
		assertEq(t, s2.Len(), 0)
		assertEq(t, s3.Len(), 0)
	})

	t.Run("sets", func(t *testing.T) {
		s1 := New([]int{1, 2, 3, 4})
		s2 := New([]int{3, 4, 5, 6})
		assert(t, s1.Union(s2).Equal(s2.Union(s1)))
		s3 := s1.Union(s2)
		assertEq(t, s3.Len(), 6)
		for i := 1; i <= 6; i++ {
			assert(t, s3.Contains(i))
		}
		assertEq(t, s1.Len(), 4)
		assertEq(t, s2.Len(), 4)
	})
}

func TestSetDiff(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})

	// subset
	assertEq(t, s1.Diff(s2).Len(), 0)

	// superset
	diff := s2.Diff(s1)
	assertEq(t, diff.Len(), 2)
	assert(t, diff.Contains(4))
	assert(t, diff.Contains(5))

	// equal
	assertEq(t, s1.Diff(s1).Len(), 0)
}

func TestSetSymDiff(t *testing.T) {
	s1 := New([]int{0, 1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})
	symdiff := s1.SymDiff(s2)
	assertEq(t, symdiff.Len(), 3)
	assert(t, symdiff.Contains(0))
	assert(t, symdiff.Contains(4))
	assert(t, symdiff.Contains(5))
}

func TestSetIntersect(t *testing.T) {
	s1 := New([]int{0, 1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})
	intersect := s1.Intersect(s2)
	assertEq(t, intersect.Len(), 3)
	assert(t, intersect.Contains(1))
	assert(t, intersect.Contains(2))
	assert(t, intersect.Contains(3))

	// no intersection
	s1 = New([]int{1, 2, 3})
	s2 = New([]int{4, 5})
	intersect = s1.Intersect(s2)
	assertEq(t, intersect.Len(), 0)
}

func TestSetDisjoint(t *testing.T) {
	s1 := New([]int{0, 1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})
	s3 := New([]int{4, 5})
	s4 := New[int](nil)
	assert(t, s1.Disjoint(s3))
	assert(t, !s1.Disjoint(s2))
	// set theory considers all empty sets to be unique
	assert(t, s4.Disjoint(s4))
}

func TestSetSubset(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})
	s3 := New([]int{0, 1, 2, 3})
	s4 := New[int](nil)
	assert(t, s4.Subset(s1))
	assert(t, s1.Subset(s2))
	assert(t, !s2.Subset(s1))
	assert(t, !s3.Subset(s2))
	assert(t, s4.Subset(s4))
}

func TestSetSuperset(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s2 := New([]int{1, 2, 3, 4, 5})
	s3 := New([]int{0, 1, 2, 3})
	s4 := New[int](nil)
	assert(t, s2.Superset(s1))
	assert(t, !s1.Superset(s2))
	assert(t, !s2.Superset(s3))
	assert(t, s4.Superset(s4))
}

func TestSetString(t *testing.T) {
	assertEq(t, New[int](nil).String(), "{}")
	assertEq(t, New([]int{1}).String(), "{1}")

	// sets aren't ordered so check any combination
	str := New([]int{1, 2}).String()
	if str != "{1,2}" && str != "{2,1}" {
		t.Errorf("unexpected String(): got %q, want one of: [{1,2},{2,1}]", str)
	}
}

// intentionally does not use set since it's used for testing the set
func sliceContains[T comparable](elems []T, el T) bool {
	for _, e := range elems {
		if e == el {
			return true
		}
	}
	return false
}

func assertEq[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if want != got {
		t.Errorf("assertion failed, want %v, got %v", want, got)
	}
}

func assert(t *testing.T, pass bool) {
	t.Helper()
	if !pass {
		t.Errorf("assertion failed")
	}
}
