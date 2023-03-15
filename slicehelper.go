package slicehelper

import (
	"errors"
	"reflect"
	"unsafe"
)

// Reverse reverses the elements of 's' in place returning the same but modified slice.
func Reverse[S ~[]E, E any](s S) S {
	if len(s) > 1 {
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		}
	}
	return s
}

// ReverseNew returns the new slice contining the reversed elements of 's'.
func ReverseNew[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	if len(s) == 0 {
		return S{}
	}
	s2 := make(S, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		s2 = append(s2, s[i])
	}
	return s2
}

// RemoveInPlace removes in place from 's' the element at the specified index.
func RemoveInPlace[S ~[]E, E any](s S, idx int) (S, error) {
	if len(s) == 0 {
		return nil, errors.New("nil or empty slice")
	}
	if !(0 <= idx && idx < len(s)) {
		return nil, errors.New("wrong index")
	}
	copy(s[idx:], s[idx+1:])
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Len = len(s) - 1
	return s, nil
}

// Split splits a sequence of ints [0..'len'-1] (indexes of a slice with length 'len') into 'n' (equal as possible) parts.
//
// 'len' and 'n' must be greater than 1. 'len' must not be less than 'n'.
// The result contains start indexes of the parts and 'len'.
// So each consecutive pair of result's elements may be used as 'low' and 'high' indices for subslicing.
//
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
	var ii []int
	remainder := len
	partStart := 0
	for remainder > 0 {
		ii = append(ii, partStart)
		partLen := remainder / n
		if remainder%n > 0 {
			partLen++
		}
		partStart += partLen
		remainder -= partLen
		n--
	}
	ii = append(ii, len)
	return ii, nil
}
