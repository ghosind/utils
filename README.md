# Utilities for Golang

![Test](https://github.com/ghosind/utils/workflows/utils/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/utils/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/utils)
[![Latest version](https://img.shields.io/github/v/release/ghosind/utils?include_prereleases)](https://github.com/ghosind/utils)
![License Badge](https://img.shields.io/github/license/ghosind/utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/utils.svg)](https://pkg.go.dev/github.com/ghosind/utils)

A set of utilities functions for Golang.

> PLEASE NOTE: This package is working with Go 1.18 and later versions.

## Installation

You can install this package with [`go get`](https://golang.org/cmd/go) command:

```sh
go get -u github.com/ghosind/utils
```

## APIs

### Conditional

- `Conditional[T any](cond bool, trueExpr, falseExpr T) T`

  An alternative of ternary operator, same of `cond ? trueExpr : falseExpr`.

- `Max[T constraints.Ordered](a, b T) T`

  Gets the maximum value of two values.

- `Min[T constraints.Ordered](a, b T) T`

  Gets the minimum value of two values.

### Type

- `IsSameType(v1, v2 any) bool`

  Compares two values' type.

- `IsSameRawType(v1, v2 any) bool`

  Compares two values' type without pointer.

- `TypeOf(v any) string`

  Gets the type of the value represented in string.

- `RawTypeOf(v any) string`

  Gets the type string name without pointer.

- `GetElem(o any) any`

  Gets element without pointer.

### Pointer and Value

- `Pointer[T any](v T) *T`

  Gets the pointer of a value.

- `func Value[T any](v *T) T`

  Gets the value of a pointer, or the zero value of the type if the pointer is nil.

- `ValueWithDefault[T any](v *T, defaultValue T) T`

  Gets the value of a pointer, or the default value if the pointer is nil.

- `PointerSlice[T any](v []T) *[]T`

  Converts a slice to a pointer slice.

- `ValueSlice[T any](v *[]T) []T`

  Converts a pointer slice to a slice.

- `PointerMap[K comparable, V any](v map[K]V) map[K]*V`

  Converts a map to a pointer map.

- `ValueMap[K comparable, V any](v map[K]*V) map[K]V`

  Converts a pointer map to a map.

## License

Distributed under the MIT License. See LICENSE file for more information.
