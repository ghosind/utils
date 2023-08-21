package utils

import (
	"reflect"
)

// IsComparableType gets the value type and checks whether the type is comparable or not.
//
//	IsComparableType(1) // true for int
//	IsComparableType(1.0) // true for float64
//	IsComparableType("") // true for string
//	IsComparableType(nil) // false for nil
//	IsComparableType(SomeStruct{}) // false for struct
//	IsComparableType(&SomeStruct{}) // false for pointer
func IsComparableType(v any) bool {
	switch v.(type) {
	case
		int, int8, int16, int32, int64, // Signed integer
		uint, uint8, uint16, uint32, uint64, uintptr, // Unsigned integer
		float32, float64, // Floating-point number
		string: // string
		return true
	default:
		return false
	}
}

// IsSameType compares two values' type.
//
//	v1 := 1 // int
//	v2 := 2 // int
//	v3 := Pointer(3) // *int
//	IsSameType(v1, v2) // true
//	IsSameType(v1, v3) // false
func IsSameType(v1, v2 any) bool {
	ty1 := TypeOf(v1)
	ty2 := TypeOf(v2)

	return ty1 == ty2
}

// IsSameRawType compares two values' type without pointer.
//
//	v1 := 1 // int
//	v2 := 2 // int
//	v3 := Pointer(3) // *int
//	IsSameRawType(v1, v2) // true
//	IsSameRawType(v1, v3) // true
func IsSameRawType(v1, v2 any) bool {
	return RawTypeOf(v1) == RawTypeOf(v2)
}

// TypeOf returns the type of the value represented in string.
//
//	TypeOf(nil) // "<nil>"
//	TypeOf(1) // "int"
//	TypeOf("test") // "string"
//	TypeOf(&http.Client{}) // "*http.Client"
func TypeOf(v any) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return "<nil>"
	}

	return t.String()
}

// RawTypeOf returns the type string name without pointer.
//
//	RawTypeOf(nil) // "<nil>"
//	RawTypeOf(1) // "int"
//	RawTypeOf("test") // "string"
//	RawTypeOf(&http.Client{}) // "http.Client"
func RawTypeOf(v any) string {
	ty := reflect.TypeOf(v)

	if ty == nil {
		return "<nil>"
	}

	for ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	return ty.String()
}

// GetElem returns the element without pointer.
//
//	v := 1
//	vp := &v
//	GetElem(v) // 1
//	GetElem(vp) // 1
func GetElem(o any) any {
	if o == nil {
		return nil
	}

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.Interface()
}
