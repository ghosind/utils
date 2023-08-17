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

// Min returns the minimum value between v1 and v2.
//
//	Min(1, 2) // 1
//	Min(1.0, 2.0) // 1.0
//	Min("str1", "str2") // "str1"
func Min[T constraints.Ordered](v1, v2 T) T {
	return Conditional(v1 < v2, v1, v2)
}
