package utils

// Pointer returns a pointer to the value passed in.
func Pointer[T any](v T) *T {
	return &v
}

// ValueWithDefault returns the value passed in if it is not nil, otherwise returns the default value.
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
