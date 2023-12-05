package utils

import "errors"

var (
	// ErrNilDest means the destination variable is empty.
	ErrNilDest error = errors.New("destination is nil")

	// ErrNilFunc means the function from parameter is nil.
	ErrNilFunc error = errors.New("function is nil")
)
