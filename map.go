package utils

// CloneMap creates a new map, and copies all the keys and their value from the source map into
// the new map. It will return nil if the source map is nil.
func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}

	dest := make(map[K]V, len(src))
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

// CopyMap copies all keys and their value from the source map into the destination map, and it
// will overwrite the old value in the destination map if the key has existed. The function will
// return a nil destination error if the source map is not nil and the destination map is nil.
func CopyMap[K comparable, V any](dest, src map[K]V) error {
	if src != nil && dest == nil {
		return ErrNilDest
	}

	for k, v := range src {
		dest[k] = v
	}

	return nil
}
