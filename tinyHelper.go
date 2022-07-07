package tinyHelper

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	prfBase = "0"
	Base02  = "b"
	Base08  = "o"
	Base16  = "x"
	digits  = "0123456789abcdefghijklmnopqrstuvwxyz"
)

type IntTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ComparableTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

type NumsTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type GeneralTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128 | ~string | ~bool
}

// ToStr[GeneralTypes](GeneralTypes) string
// Wrapper for convert to string by 'fmt.Sprintf'
func ToStr[T GeneralTypes](val T) (nv string) {
	return fmt.Sprintf("%v", val)
}

// toBaseOther(int, int) string
// Wrapper for convert to another base
func toBaseOther(v, base int) string {
	if v/base > 0 {
		return toBaseOther(v/base, base) + string(digits[v%base])
	} else {
		return string(digits[v%base])
	}
}

// ToBase[GeneralTypes](T, int, int, ...int) string
// Wrapper for convert from one base to another base
func ToBase[T GeneralTypes](val T, from, to int, prefixNulls ...int) (nv string) {
    to10 := ToInt(val, from)
	nv = ToStr(to10)
	prx := ""
	if prefixNulls != nil || len(prefixNulls) > 0 {
		prx = ToStr(prefixNulls[0])
	}
	switch to {
	case 2:
		nv = fmt.Sprintf("%"+prfBase+prx+Base02, to10)
	case 8:
		nv = fmt.Sprintf("%"+prfBase+prx+Base08, to10)
	case 16:
		nv = fmt.Sprintf("%"+prfBase+prx+Base16, to10)
	default:
        nv = fmt.Sprintf("%"+prfBase+prx+"v", toBaseOther(to10, to))
	}
	return
}

// ToNum[NumsTypes,GeneralTypes](GeneralTypes) NumsTypes {
// Wrapper for convert into numerable types.
// !!! Caution, if val more than T1 bitSize - may return quotient of division by dimension
// Using example: 'ToNum[int](true)'
//                'ToNum[float64](float64(122.1)) + 1)'
func ToNum[T1 NumsTypes, T GeneralTypes](val T, base ...int) T1 {
	f := ToStr(val)
	parBase := 10 // decimal as default
	if base != nil || len(base) > 0 {
		parBase = base[0] // tests result max == 36 (26 alphabet + 10 digits), more == 10 base
	}
	// Del prefix. Parse values with prefix 0b, 0o, 0x works if base '0' value in parameter.
	if len(f) >= 2 {
		switch {
        case f[:2] == (prfBase + Base02) && parBase == 16:
        case f[:2] == (prfBase + Base02):
            parBase = 2
			f = f[2:]
		case f[:2] == (prfBase + Base08):
            parBase = 8
			f = f[2:]
		case f[:2] == (prfBase + Base16):
            parBase = 16
			f = f[2:]
		}
	}

	if nv, err := strconv.ParseInt(f, parBase, strconv.IntSize); err == nil {
		return T1(nv)
	}
	if nv, err := strconv.ParseUint(f, parBase, strconv.IntSize); err == nil {
		return T1(nv)
	}
	if nv, err := strconv.ParseFloat(f, strconv.IntSize); err == nil {
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

// ToInt[GeneralTypes](GeneralTypes, ...int) int {
// Wrapper for convert into 'int'.
func ToInt[T GeneralTypes](val T, base ...int) int {
	return ToNum[int](val, base...)
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
