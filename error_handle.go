package utils

import (
	"fmt"
)

// TryCatch is an alternative to the `try...catch...finally` statement. This function requires a
// `try` statement function and a `catch` statement function, and it accepts an optional `finally`
// statement function.
//
// This function will execute the try statement function, and call the catch statement function if
// the try function returns an error or panic. It will always perform the `finally` statement
// function whether the try function is successful or not if the `finally` function is not nil.
// However, if the `catch` function panics, this TryCatch statement will also panic, and finally
// function will never run.
//
//	err := TryCatch(
//	  func() error {
//	    // Do something as try statement
//	    return nil
//	  },
//	  func (err error) {
//	    // Do something as catch statement
//	  },
//	  func () {
//	    // An optional finally statement
//	  },
//	)
func TryCatch(
	tryStmt func() error,
	catchStmt func(error),
	finallyStmt ...func(),
) error {
	var tryError error
	func() {
		defer func() {
			if pe := recover(); pe != nil {
				if e, ok := pe.(error); ok {
					tryError = e
				} else {
					tryError = fmt.Errorf("%v", pe)
				}
			}
		}()

		if tryStmt != nil {
			tryError = tryStmt()
		}
	}()

	if tryError != nil && catchStmt != nil {
		catchStmt(tryError)
	}

	if len(finallyStmt) > 0 && finallyStmt[0] != nil {
		finallyStmt[0]()
	}

	return tryError
}
