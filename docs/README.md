# slicehelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/slicehelper.svg)](https://pkg.go.dev/github.com/solsw/slicehelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/slicehelper)

Generic helpers for Go's [slice](https://go.dev/ref/spec#Slice_types).

## Installation

```sh
go get github.com/solsw/slicehelper
```

## Result and the input slice

The functions differ in how their result relates to the input slice:

| Behaviour | Functions |
| --- | --- |
| Always return a newly allocated slice | `ReverseNew`, `Project` |
| Return a new slice, or the input unchanged when there is nothing to do | `Filter` |
| Return a subslice sharing the input's backing array | `TrimStart`, `TrimEnd`, `Trim` |
| Modify the input in place | `RemoveInPlace`, `SortDesc` |

## Functions

### ReverseNew

```go
func ReverseNew[S ~[]E, E any](s S) S
```

Returns a new slice containing the reversed elements of `s`. The initial slice
remains unchanged. If `s` is nil, nil is returned.

```go
slicehelper.ReverseNew([]int{1, 2, 3}) // [3 2 1]
```

### SortDesc

```go
func SortDesc[S ~[]E, E cmp.Ordered](x S)
```

Like [`slices.Sort`](https://pkg.go.dev/slices#Sort) but sorts `x` in
descending order. `x` is sorted in place.

```go
x := []int{3, 1, 2}
slicehelper.SortDesc(x) // x == [3 2 1]
```

### TrimStart / TrimEnd / Trim

```go
func TrimStart[S ~[]E, E any](s S, pred func(E) bool) S
func TrimEnd[S ~[]E, E any](s S, pred func(E) bool) S
func Trim[S ~[]E, E any](s S, pred func(E) bool) S
```

Remove the leading (`TrimStart`), trailing (`TrimEnd`), or both leading and
trailing (`Trim`) elements of `s` that satisfy `pred`. The result is a subslice
that shares `s` backing array. If `s` is nil or empty or `pred` is nil, `s` is
returned.

```go
blank := func(s string) bool { return s == "" }
slicehelper.Trim([]string{"", "a", "", "b", ""}, blank) // [a "" b]
```

### Split

```go
func Split(length, n int) ([]int, error)
```

Splits a sequence of ints `[0..length-1]` (the indexes of a slice of length
`length`) into `n` as-equal-as-possible integer parts. `length` and `n` must be
greater than 1, and `length` must not be less than `n`. The result holds the
start indexes of the parts plus `length`, so each consecutive pair can be used
as `low` and `high` indices for subslicing. Intended for splitting a slice for
(parallel) processing of the parts.

```go
ii, _ := slicehelper.Split(8, 3) // [0 3 6 8]
// parts: s[0:3], s[3:6], s[6:8]
```

### RemoveInPlace

```go
func RemoveInPlace[S ~[]E, E any](s S, idx int) (S, error)
```

Removes the element at index `idx` from `s`. `s` is modified in place and the
resulting (shortened) slice is returned. Returns an error for a nil/empty slice
or an out-of-range index.

```go
r, _ := slicehelper.RemoveInPlace([]int{1, 2, 3}, 1) // [1 3]
```

### Project

```go
func Project[S ~[]Es, R ~[]Er, Es, Er any](s S, proj func(Es) Er) R
```

Projects the elements of `s` into a new slice with the help of `proj`. If `s`
is nil or `proj` is nil, nil is returned. The result element type cannot be
inferred, so it must be supplied explicitly.

```go
lengths := slicehelper.Project[[]string, []int](
	[]string{"zero", "one"},
	func(s string) int { return len(s) },
) // [4 3]
```

### Filter

```go
func Filter[S ~[]E, E any](s S, pred func(E) bool) S
```

Returns a new slice containing only the elements of `s` that satisfy `pred`. If
`s` is nil, nil is returned. If `pred` is nil, `s` is returned.

```go
slicehelper.Filter([]int{1, 2, 3, 4}, func(i int) bool { return i%2 == 0 }) // [2 4]
```
