# Utilities for Golang

![Test](https://github.com/ghosind/utils/workflows/utils/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/utils/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/utils)
[![Latest version](https://img.shields.io/github/v/release/ghosind/utils?include_prereleases)](https://github.com/ghosind/utils)
![License Badge](https://img.shields.io/github/license/ghosind/utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/utils.svg)](https://pkg.go.dev/github.com/ghosind/utils)

A set of utilities functions for Golang.

> PLEASE NOTE: This package is working with Go 1.18 and later versions.

- [Installation](#installation)
- [Getting Started](#getting-started)
  - [Handle Pointers](#handle-pointers)
  - [Conditional Expression](#conditional-expression)
- [Available Utilities](#available-utilities)
  - [Conditional](#conditional)
  - [Map Manipulation](#map-manipulation)
  - [Type](#type)
  - [Pointer and Value](#pointer-and-value)
- [License](#license)

## Installation

You can install this package with [`go get`](https://golang.org/cmd/go) command:

```sh
go get -u github.com/ghosind/utils
```

## Getting Started

### Handle Pointers

With the `Pointer` method, you can easy to get a pointer that points to any value you want.

```go
str := "Hello world" // string
// get the pointer of the string, it equal to strp := &str
strp := utils.Pointer(str) // string pointer, and it point to str
log.Print(*strp) // Hello world

// It's also working for literal values
intp := utils.Pointer(1)
// You can't do:
// intp := &1
log.Print(*intp) // 1
```

You can also use `Value` or `ValueWithDefault` to get the value of the pointer. For nil pointer, `Value` method will return the zero value of the type, and `ValueWithDefault` method will return the default value you set.

```go
utils.Value(strp) // Hello world
utils.ValueWithDefault(str, "Default") // Hello world

strp = nil
utils.Value(strp) // "" (empty string)
utils.ValueWithDefault(str, "Default") // Default
```

### Conditional Expression

Golang does not provided ternary operator (`?:`), but you can use `utils.Conditional` or `utils.ConditionalExpr` to make it like a ternary expression.

```go
a := 1
b := 2
bigger := utils.Conditional(a > b, a, b) // a > b ? a : b
// bigger: 2
```

## Available Utilities

### Conditional

- [`Conditional`](https://pkg.go.dev/github.com/ghosind/utils#Conditional): An alternative of ternary operator, same of `cond ? trueExpr : falseExpr`.

- [`ConditionalExpr`](https://pkg.go.dev/github.com/ghosind/utils#ConditionalExpr): An alternative to the conditional (ternary) operator (?:), it'll run expression by the conditional result.

- [`Max`](https://pkg.go.dev/github.com/ghosind/utils#Max): Gets the maximum value between the two values.

- [`MaxN`](https://pkg.go.dev/github.com/ghosind/utils#MaxN): Returns the maximum value in the list and returns the zero value of the type if no parameter.

- [`Min`](https://pkg.go.dev/github.com/ghosind/utils#Min): Gets the minimum value between the two values.

- [`MinN`](https://pkg.go.dev/github.com/ghosind/utils#MinN): Returns the minimum value in the list and returns the zero value of the type if no parameter.

### Error Handling

- [`TryCatch`](https://pkg.go.dev/github.com/ghosind/utils#TryCatch): An alternative of `try...catch...finally` statement.

### Math

- [`Matrix`](https://pkg.go.dev/github.com/ghosind/utils#Matrix): Creates and initializes a n*m matrix.

### Map Manipulation

- [`CloneMap`](https://pkg.go.dev/github.com/ghosind/utils#CloneMap): Creates a new map, and copies all the keys and their value from the source map.

- [`CopyMap`](https://pkg.go.dev/github.com/ghosind/utils#CopyMap): Copies all keys and their value from the source map into the destination map.

### Type

- [`IsComparableType`](https://pkg.go.dev/github.com/ghosind/utils#IsComparableType): Check the type of value is comparable or not.

- [`IsSameType`](https://pkg.go.dev/github.com/ghosind/utils#IsSameType): Compares two values' type.

- [`IsSameRawType`](https://pkg.go.dev/github.com/ghosind/utils#IsSameRawType): Compares two values' type without pointer.

- [`TypeOf`](https://pkg.go.dev/github.com/ghosind/utils#TypeOf): Gets the type of the value represented in string.

- [`RawTypeOf`](https://pkg.go.dev/github.com/ghosind/utils#RawTypeOf): Gets the type string name without pointer.

- [`GetElem`](https://pkg.go.dev/github.com/ghosind/utils#GetElem): Gets element without pointer.

### Pointer and Value

- [`Pointer`](https://pkg.go.dev/github.com/ghosind/utils#Pointer): Gets the pointer of a value.

- [`Value`](https://pkg.go.dev/github.com/ghosind/utils#Value): Gets the value of a pointer, or the zero value of the type if the pointer is nil.

- [`ValueWithDefault`](https://pkg.go.dev/github.com/ghosind/utils#ValueWithDefault): Gets the value of a pointer, or the default value if the pointer is nil.

- [`PointerSlice`](https://pkg.go.dev/github.com/ghosind/utils#PointerSlice): Converts a slice to a pointer slice.

- [`ValueSlice`](https://pkg.go.dev/github.com/ghosind/utils#ValueSlice): Converts a pointer slice to a slice.

- [`PointerMap`](https://pkg.go.dev/github.com/ghosind/utils#PointerMap): Converts a map to a pointer map.

- [`ValueMap`](https://pkg.go.dev/github.com/ghosind/utils#ValueMap): Converts a pointer map to a map.

## License

Distributed under the MIT License. See LICENSE file for more information.
