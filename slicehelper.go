package slicehelper

import (
	"cmp"
	"errors"
	"slices"

	"github.com/solsw/mathhelper"
)

// ReverseNew returns new slice containing the reversed elements of 's'.
// The initial slice remains unchanged. If 's' is nil, nil is returned.
func ReverseNew[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	s2 := make(S, len(s))
	copy(s2, s)
	slices.Reverse(s2)
	return s2
}

// SortDesc is like [slices.Sort] but sorts a slice in descending order.
// 'x' is sorted in place.
func SortDesc[S ~[]E, E cmp.Ordered](x S) {
	slices.SortFunc(x, func(a, b E) int {
		return cmp.Compare(b, a)
	})
}

// TrimStart removes all the leading elements from 's' that satisfy 'pred'.
// The result is a subslice that shares 's' backing array.
// If 's' is nil or empty or 'pred' is nil, 's' is returned.
func TrimStart[S ~[]E, E any](s S, pred func(E) bool) S {
	if len(s) == 0 || pred == nil {
		return s
	}
	beg := 0
	for beg < len(s) && pred(s[beg]) {
		beg++
	}
	return s[beg:]
}

// TrimEnd removes all the elements from 's' that satisfy 'pred'.
// The result is a subslice that shares 's' backing array.
// If 's' is nil or empty or 'pred' is nil, 's' is returned.
func TrimEnd[S ~[]E, E any](s S, pred func(E) bool) S {
	if len(s) == 0 || pred == nil {
		return s
	}
	end := len(s)
	for end > 0 && pred(s[end-1]) {
		end--
	}
	return s[:end]
}

// Trim removes all leading and trailing elements from 's' that satisfy 'pred'.
// The result is a subslice that shares 's' backing array.
// If 's' is nil or empty or 'pred' is nil, 's' is returned.
func Trim[S ~[]E, E any](s S, pred func(E) bool) S {
	return TrimEnd(TrimStart(s, pred), pred)
}

// Split splits a sequence of ints [0..'length'-1] (indexes of a slice
// with length 'length') into 'n' as equal as possible integer parts.
// 'length' and 'n' must be greater than 1. 'length' must not be less than 'n'.
// The result contains start indexes of the parts and 'length', so each consecutive pair
// of the result's elements may be used as 'low' and 'high' indices for subslicing
// (see TestSplit for example).
// Function is intended for splitting a slice for (parallel) processing of the parts.
func Split(length, n int) ([]int, error) {
	if length <= 1 {
		return nil, errors.New("wrong length")
	}
	if n <= 1 {
		return nil, errors.New("wrong parts number")
	}
	if length < n {
		return nil, errors.New("length less than parts number")
	}
	ii := []int{0}
	last := 0
	pp, _ := mathhelper.Split(length, n)
	for _, p := range pp {
		last += p
		ii = append(ii, last)
	}
	return ii, nil
}

// RemoveInPlace removes in place from 's' the element at the specified index.
// 's' is modified in place and the resulting (shortened) slice is returned.
func RemoveInPlace[S ~[]E, E any](s S, idx int) (S, error) {
	if s == nil {
		return nil, errors.New("nil slice")
	}
	if len(s) == 0 {
		return nil, errors.New("empty slice")
	}
	if !(0 <= idx && idx < len(s)) {
		return nil, errors.New("wrong index")
	}
	return slices.Delete(s, idx, idx+1), nil
}

// Project projects the elements of 's' into a new slice with the help of 'proj'.
// If 's' is nil or 'proj' is nil, nil is returned.
func Project[S ~[]Es, R ~[]Er, Es, Er any](s S, proj func(Es) Er) R {
	if s == nil || proj == nil {
		return nil
	}
	rr := make(R, 0, len(s))
	for _, es := range s {
		rr = append(rr, proj(es))
	}
	return rr
}

// Filter returns a new slice containing only the elements of 's' that satisfy 'pred'.
// If 's' is nil, nil is returned.
// If 'pred' is nil, 's' is returned.
func Filter[S ~[]E, E any](s S, pred func(E) bool) S {
	if s == nil {
		return nil
	}
	if len(s) == 0 || pred == nil {
		return s
	}
	ss := make(S, 0, len(s))
	for _, e := range s {
		if pred(e) {
			ss = append(ss, e)
		}
	}
	return ss
}
