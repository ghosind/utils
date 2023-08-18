package utils

// Pointer returns a pointer to the value passed in.
//
//	v := 1 // type is int
//	vp := Pointer(v) // type is *int, *vp = 1
//	TypeOf(vp) // "*int"
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

// ValueSlice converts a slice of pointers into a slice of values, and it'll
// set the value to the zero value of the type if the pointer is point to nil.
//
//	s1 := []*string{Pointer("Hello"), nil, Pointer("World")}
//	s2 := ValueSlice(s1)
//	// ["Hello", "", "World"]
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

// ValueMap converts a map of pointers into a map of values, and it'll set the
// value to the zero value of the type if the pointer is point to nil.
//
//	m1 := map[string]*int{ "a": Pointer(1), "b": nil, "c": Pointer(3) }
//	m2 := ValueMap(m1)
//	// [a:1, b:0, c:3]
func ValueMap[K comparable, V any](src map[K]*V) map[K]V {
	dest := make(map[K]V)
	for k, val := range src {
		if val != nil {
			dest[k] = *val
		}
	}

	return dest
}
