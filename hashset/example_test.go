package hashset_test

import (
	"fmt"
	"sort"

	"github.com/scnewma/go-sets/hashset"
)

func Example() {
	// hash set with known list of items
	_ = hashset.New([]string{"Einar", "Olaf", "Harald"})

	// empty hash set
	books := hashset.New[string](nil)
	books.Insert("A Dance With Dragons")
	books.Insert("To Kill a Mockingbird")
	books.Insert("The Odyssey")
	books.Insert("The Great Gatsby")

	// Check for a specific book
	if !books.Contains("The Winds of Winter") {
		fmt.Printf("We have %d books, but that isn't one\n", books.Len())
	}

	// remove a book
	books.Remove("The Odyssey")

	// print them out (sorted bc order is arbitrary)
	bs := books.Elems()
	sort.Strings(bs)
	for _, b := range bs {
		fmt.Println(b)
	}

	// Output:
	// We have 4 books, but that isn't one
	// A Dance With Dragons
	// The Great Gatsby
	// To Kill a Mockingbird
}

func ExampleSet_Disjoint() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New[int](nil)

	fmt.Println(a.Disjoint(b))
	b.Insert(4)
	fmt.Println(a.Disjoint(b))
	b.Insert(1)
	fmt.Println(a.Disjoint(b))

	// Output:
	// true
	// true
	// false
}

func ExampleSet_Intersect() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	// {2, 3} in arbitrary order
	i := a.Intersect(b)
	fmt.Println(i.Contains(2))
	fmt.Println(i.Contains(3))

	// Output:
	// true
	// true
}

func ExampleSet_Union() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	// {1, 2, 3, 4} in arbitrary order
	i := a.Union(b)
	fmt.Println(i.Contains(1))
	fmt.Println(i.Contains(2))
	fmt.Println(i.Contains(3))
	fmt.Println(i.Contains(4))

	// Output:
	// true
	// true
	// true
	// true
}

func ExampleSet_Contains() {
	set := hashset.New([]int{1, 2, 3})

	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(4))

	// Output:
	// true
	// false
}

func ExampleSet_Subset() {
	sup := hashset.New([]int{1, 2, 3})
	set := hashset.New[int](nil)

	fmt.Println(set.Subset(sup))
	set.Insert(2)
	fmt.Println(set.Subset(sup))
	set.Insert(4)
	fmt.Println(set.Subset(sup))

	// Output:
	// true
	// true
	// false
}

func ExampleSet_Superset() {
	sub := hashset.New([]int{1, 2})
	set := hashset.New[int](nil)

	fmt.Println(set.Superset(sub))

	set.Insert(0)
	set.Insert(1)
	fmt.Println(set.Superset(sub))

	set.Insert(2)
	fmt.Println(set.Superset(sub))

	// Output:
	// false
	// false
	// true
}

func ExampleSet_Insert() {
	set := hashset.New[int](nil)

	fmt.Println(set.Insert(2))
	fmt.Println(set.Insert(2))
	fmt.Println(set.Len())

	// Output:
	// true
	// false
	// 1
}

func ExampleSet_Remove() {
	set := hashset.New[int](nil)

	set.Insert(2)
	fmt.Println(set.Remove(2))
	fmt.Println(set.Remove(2))

	// Output:
	// true
	// false
}

func ExampleSet_Merge() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	fmt.Println(a.Len())
	fmt.Println(b.Len())
	a.Merge(b)
	fmt.Println(a.Len())
	fmt.Println(b.Len())

	// Output:
	// 3
	// 3
	// 4
	// 3
}

func ExampleSet_Diff() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	fmt.Println(a.Diff(b).String())

	// Note that Diff is not symmetric
	fmt.Println(b.Diff(a).String())

	// Output:
	// {1}
	// {4}
}

func ExampleSet_SymDiff() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	// {1, 4} in arbitrary order
	d := a.SymDiff(b)
	fmt.Println(d.Contains(1))
	fmt.Println(d.Contains(4))

	ab := a.SymDiff(b)
	ba := b.SymDiff(a)
	fmt.Println(ab.Equal(ba))

	// Output:
	// true
	// true
	// true
}

func ExampleSet_Len() {
	set := hashset.New[int](nil)

	fmt.Println(set.Len())
	set.Insert(1)
	fmt.Println(set.Len())

	// Output:
	// 0
	// 1
}

func ExampleSet_Clear() {
	set := hashset.New[int](nil)
	set.Insert(1)
	set.Clear()
	fmt.Println(set.Len())

	// Output:
	// 0
}

func ExampleSet_Equal() {
	a := hashset.New([]int{1, 2, 3})
	b := hashset.New([]int{4, 2, 3, 4})

	fmt.Println(a.Equal(b))
	a.Insert(4)
	fmt.Println(a.Equal(b))
	b.Insert(1)
	fmt.Println(a.Equal(b))

	// Output:
	// false
	// false
	// true
}
