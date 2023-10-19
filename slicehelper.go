package slicehelper

import (
	"cmp"
	"errors"
	"slices"
	"unsafe"

	"github.com/solsw/mathhelper"
)

// ReverseNew returns new slice containing the reversed elements of 's'.
// The initial slice remains unchanged. If 's' is nil, nil is returned.
func ReverseNew[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	if len(s) == 0 {
		return S{}
	}
	if len(s) == 1 {
		return S{s[0]}
	}
	s2 := make(S, len(s))
	copy(s2, s)
	slices.Reverse[S, E](s2)
	return s2
}

// SortDesc is like [slices.Sort] but sorts a slice in descending order.
func SortDesc[S ~[]E, E cmp.Ordered](x S) {
	slices.SortFunc[S, E](x, func(a, b E) int {
		return -cmp.Compare[E](a, b)
	})
}

// TrimStart removes all the leading elements from 's' that satisfy 'f'.
// If 's' is nil or empty or 'f' is nil, 's' is returned.
func TrimStart[S ~[]E, E any](s S, f func(E) bool) S {
	if len(s) == 0 || f == nil {
		return s
	}
	beg := 0
	for beg < len(s) && f(s[beg]) {
		beg++
	}
	return s[beg:]
}

// TrimEnd removes all the elements from 's' that satisfy 'f'.
// If 's' is nil or empty or 'f' is nil, 's' is returned.
func TrimEnd[S ~[]E, E any](s S, f func(E) bool) S {
	if len(s) == 0 || f == nil {
		return s
	}
	end := len(s)
	for end > 0 && f(s[end-1]) {
		end--
	}
	return s[:end]
}

// Trim removes all leading and trailing elements from 's' that satisfy 'f'.
// If 's' is nil or empty or 'f' is nil, 's' is returned.
func Trim[S ~[]E, E any](s S, f func(E) bool) S {
	return TrimEnd[S, E](TrimStart[S, E](s, f), f)
}

// Split splits a sequence of ints [0..'len'-1] (indexes of a slice with length 'len')
// into 'n' as equal as possible integer parts.
// 'len' and 'n' must be greater than 1. 'len' must not be less than 'n'.
// The result contains start indexes of the parts and 'len',
// so each consecutive pair of result's elements may be used as 'low' and 'high' indices for subslicing.
// Function is intended for splitting a slice or an array for (parallel) processing of the parts.
func Split(len, n int) ([]int, error) {
	if len <= 1 {
		return nil, errors.New("wrong length")
	}
	if n <= 1 {
		return nil, errors.New("wrong parts number")
	}
	if len < n {
		return nil, errors.New("length less than parts number")
	}
	ii := []int{0}
	last := 0
	pp, _ := mathhelper.Split(len, n)
	for _, p := range pp {
		last += p
		ii = append(ii, last)
	}
	return ii, nil
}

// RemoveInPlace removes in place from 's' the element at the specified index.
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
	copy(s[idx:], s[idx+1:])
	return unsafe.Slice(&s[0], len(s)-1), nil
}
