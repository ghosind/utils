package utils

import "golang.org/x/exp/constraints"

// Conditional is an alternative to the conditional (ternary) operator (?:).
func Conditional[T any](cond bool, trueExpr, falseExpr T) T {
	if cond {
		return trueExpr
	} else {
		return falseExpr
	}
}

// ConditionalExpr is an alternative to the conditional (ternary) operator (?:), it'll run
// expression by the conditional result.
func ConditionalExpr[T any](cond bool, trueExpr, falseExpr func() T) T {
	if cond {
		return trueExpr()
	} else {
		return falseExpr()
	}
}

// Max returns the maximum value between v1 and v2.
func Max[T constraints.Ordered](v1, v2 T) T {
	return Conditional(v1 > v2, v1, v2)
}

// Min returns the minimum value between v1 and v2.
func Min[T constraints.Ordered](v1, v2 T) T {
	return Conditional(v1 < v2, v1, v2)
}
