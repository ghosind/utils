package utils

import (
	"errors"
	"fmt"
)

// Try provided a recoverable function container to handle both return error and panic error.
//
//	Try(func() error {
//	  panic("some error")
//	})
//	// some error
func Try(fn func() error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch t := e.(type) {
			case error:
				err = t
			case string:
				err = errors.New(t)
			default:
				err = fmt.Errorf("%v", t)
			}
		}
	}()

	if fn == nil {
		panic(ErrNilFunc)
	}

	err = fn()
	return
}

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
	err := Try(tryStmt)

	if err != nil && catchStmt != nil {
		catchStmt(err)
	}

	if len(finallyStmt) > 0 && finallyStmt[0] != nil {
		finallyStmt[0]()
	}

	return err
}
