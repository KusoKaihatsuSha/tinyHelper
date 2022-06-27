package tinyHelper

import (
	"fmt"
	"sort"
	"strconv"
)

type ComparableTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type NumsTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type GeneralTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128 | ~string | ~bool
}

// ToStr[GeneralTypes](GeneralTypes) (string)
// Wrapper for convert to string by 'fmt.Sprintf'
func ToStr[T GeneralTypes](val T) (nv string) {
	return fmt.Sprintf("%v", val)
}

// ToNum[NumsTypes,GeneralTypes](GeneralTypes) NumsTypes {
// Wrapper for convert into numerable types
// Using example: 'ToNum[int](true)'
//                'ToNum[float64](float64(122.1)) + 1)'
func ToNum[T1 NumsTypes, T GeneralTypes](val T) T1 {
	f := ToStr(val)
	if nv, err := strconv.ParseInt(f, 10, strconv.IntSize); err == nil {
		return T1(nv)
	}
	if nv, err := strconv.ParseUint(f, 10, strconv.IntSize); err == nil {
		return T1(nv)
	}
	if nv, err := strconv.ParseFloat(f, 10); err == nil {
		return T1(nv)
	}
	if nv, err := strconv.ParseBool(f); err == nil {
		if nv {
			return T1(1)
		}
		return T1(0)
	}
	return T1(0)
}

func ToInt[T GeneralTypes](val T) int {
	return ToNum[int](val)
}

// SliceRotate[any]([]T, ...bool) []any
// Rotate slice or rotate even/odd pair
func SliceRotate[T any](val []T, byEvenOdd ...bool) []T {
	newVal := []T{}
	for i := len(val) - 1; i >= 0; i = i - 1 - i%2 {
		pair := []T{val[i]}
		if i%2 > 0 {
			pair = append(pair, val[1^i])
		}
		if byEvenOdd != nil && byEvenOdd[0] {
			newVal = append(pair, newVal...)
		} else {
			newVal = append(newVal, pair...)
		}
	}
	return newVal
}

// SliceSort[ComparableTypes]([]ComparableTypes)
// Wrapper for sorting slice
func SliceSort[T ComparableTypes](val []T) {
	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})
}
