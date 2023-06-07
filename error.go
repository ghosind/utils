package utils

import "errors"

var (
	// ErrNilDest means the destination variable is empty.
	ErrNilDest error = errors.New("destination is nil")
)
