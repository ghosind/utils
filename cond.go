package utils

import "golang.org/x/exp/constraints"

// Conditional is an alternative to the conditional (ternary) operator (?:).
//
//	a := 1
//	b := 2
//	Conditional(a > b, a, b) // 2
func Conditional[T any](cond bool, trueExpr, falseExpr T) T {
	if cond {
		return trueExpr
	} else {
		return falseExpr
	}
}

// ConditionalExpr is an alternative to the conditional (ternary) operator (?:), it'll run
// expression by the conditional result.
//
//	var strp *string // nil
//	ConditionalExpr(strp == nil,
//	  func() string { return "" },
//	  func() string { return *strp },
//	) // ""
//
//	strp = Pointer("hello")
//	ConditionalExpr(strp == nil,
//	  func() string { return "" },
//	  func() string { return *strp },
//	) // "hello"
func ConditionalExpr[T any](cond bool, trueExpr, falseExpr func() T) T {
	if cond {
		return trueExpr()
	} else {
		return falseExpr()
	}
}

// Max returns the maximum value between v1 and v2.
//
//	Max(1, 2) // 2
//	Max(1.0, 2.0) // 2.0
//	Max("str1", "str2") // "str2"
func Max[T constraints.Ordered](v1, v2 T) T {
	return Conditional(v1 > v2, v1, v2)
}

// `MaxN` returns the maximum value in the list and returns the zero value of the type if no
// parameter.
//
//	MaxN[int]() // 0
//	MaxN(2, 1) // 2
//	MaxN(3, 2, 1, 0, 2, 4) // 4
func MaxN[T constraints.Ordered](values ...T) T {
	var max T
	if len(values) == 0 {
		return max
	}

	max = values[0]
	for i := 1; i < len(values); i++ {
		if max < values[i] {
			max = values[i]
		}
	}

	return max
}

// Min returns the minimum value between v1 and v2.
//
//	Min(1, 2) // 1
//	Min(1.0, 2.0) // 1.0
//	Min("str1", "str2") // "str1"
func Min[T constraints.Ordered](v1, v2 T) T {
	return Conditional(v1 < v2, v1, v2)
}

// `MinN` returns the minimum value in the list and returns the zero value of the type if no
// parameter.
//
//	MinN[int]() // 0
//	MinN(1, 2) // 1
//	MinN(3, 2, 1, 0, 2, 4) // 0
func MinN[T constraints.Ordered](values ...T) T {
	var min T
	if len(values) == 0 {
		return min
	}

	min = values[0]
	for i := 1; i < len(values); i++ {
		if min > values[i] {
			min = values[i]
		}
	}

	return min
}
