package utils

import "reflect"

// IsSameType compares two values' type.
func IsSameType(v1, v2 any) bool {
	ty1 := reflect.TypeOf(v1).String()
	ty2 := reflect.TypeOf(v2).String()

	return ty1 == ty2
}

// IsSameRawType compares two values' type without pointer.
func IsSameRawType(v1, v2 any) bool {
	return RawTypeOf(v1) == RawTypeOf(v2)
}

// RawTypeOf returns the type string name without pointer.
func RawTypeOf(v any) string {
	ty := reflect.TypeOf(v)

	for ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	return ty.String()
}
