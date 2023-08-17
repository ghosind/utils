package utils

// Pointer returns a pointer to the value passed in.
func Pointer[T any](v T) *T {
	return &v
}

// Value returns the value of a pointer, or the zero value of the type if the pointer is nil.
//
//	var p1 *string // nil
//	var p2 *string = Pointer("hello")
//	Value(p1) // ""
//	Value(p2) // "hello"
func Value[T any](v *T) T {
	if v != nil {
		return *v
	}

	var zero T
	return zero
}

// ValueWithDefault returns the value passed in if it is not nil, otherwise returns the default
// value.
//
//	var p1 *string // nil
//	var p2 *string = Pointer("hello")
//	ValueWithDefault(p1, "default") // "default"
//	ValueWithDefault(p2, "default") // "hello"
func ValueWithDefault[T any](val *T, defaultVal T) T {
	if val == nil {
		return defaultVal
	}
	return *val
}

// PointerToSlice converts a slice of values into a slice of pointers.
func PointerSlice[T any](src []T) []*T {
	dest := make([]*T, len(src))
	for i := 0; i < len(src); i++ {
		dest[i] = &(src[i])
	}

	return dest
}

// ValueSlice converts a slice of pointers into a slice of values.
func ValueSlice[T any](src []*T) []T {
	dest := make([]T, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dest[i] = *src[i]
		}
	}

	return dest
}

// PointerMap converts a map of values into a map of pointers.
func PointerMap[K comparable, V any](src map[K]V) map[K]*V {
	dest := make(map[K]*V)
	for k, val := range src {
		v := val
		dest[k] = &v
	}

	return dest
}

// ValueMap converts a map of pointers into a map of values.
func ValueMap[K comparable, V any](src map[K]*V) map[K]V {
	dest := make(map[K]V)
	for k, val := range src {
		if val != nil {
			dest[k] = *val
		}
	}

	return dest
}
