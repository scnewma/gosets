# go-sets

[![Go Reference](https://pkg.go.dev/badge/github.com/scnewma/gosets.svg)](https://pkg.go.dev/github.com/scnewma/gosets/hashset)

Tired of writing code like this?

```
set := map[string]struct{}{}
set["bob"] = struct{}{}

_, contains := set["bob"]
if contains {
    // ...
}
```

This package introduces a new `hashset.Set[T]` type that allows you to write this instead.

```
set := hashset.New[string](nil)
set.Insert("bob")
if set.Contains("bob") {
    // ...
}
```

In addition to basic set operations, it also supports common functions from set theory, which can be useful at times.

```
s1 := hashset.New([]string{"a", "b", "c"})
s2 := hashset.New([]string{"b", "c", "d"})
s3 := hashset.New([]string{"a"})
_ = s1.Diff(s2)      // {a}
_ = s1.SymDiff(s2)   // {a,d}
_ = s1.Intersect(s2) // {b,c}
_ = s1.Union(s2)     // {a,b,c,d}
_ = s2.Disjoint(s3)  // true
_ = s3.Subset(s1)    // true
_ = s1.Subset(s3)    // false
_ = s1.Superset(s3)  // true
```

## Acknowledgements

The API for this package was inspired by Rust's [HashSet](https://doc.rust-lang.org/std/collections/hash_set/struct.HashSet.html) implementation. Many of the examples are copies of their examples as they are simple to understand.
