package utils

import (
	"errors"
	"testing"

	"github.com/ghosind/go-assert"
)

func TestTry(t *testing.T) {
	a := assert.New(t)

	testTry(a, func() error { return nil }, false, "")
	testTry(a, func() error { return errors.New("expected error") }, true, "expected error")
	testTry(a, func() error { panic("expected panic") }, true, "expected panic")
	testTry(a, func() error { panic(errors.New("expected panic error")) }, true, "expected panic error")
	testTry(a, func() error { panic(1) }, true, "1")
}

func testTry(a *assert.Assertion, fn func() error, isError bool, errMsg string) {
	err := Try(fn)
	if !isError {
		a.NilNow(err)
	} else {
		a.NotNilNow(err)
		a.EqualNow(err.Error(), errMsg)
	}
}

func TestTryCatch(t *testing.T) {
	a := assert.New(t)

	testTryCatch(a, func() error { return nil }, true, func(err error) {})
	testTryCatch(a, func() error { return nil }, true, func(err error) {}, func() {})
	testTryCatch(a, func() error { return errors.New("test") }, false, func(err error) {})
	testTryCatch(a, func() error { return errors.New("test") }, false, func(err error) {}, func() {})
	testTryCatch(a, func() error { panic("test") }, false, func(err error) {})
	testTryCatch(a, func() error { panic("test") }, false, func(err error) {}, func() {})
	testTryCatch(a, func() error { panic(errors.New("test")) }, false, func(err error) {})
	testTryCatch(a, func() error { panic(errors.New("test")) }, false, func(err error) {}, func() {})

	testTryCatch(a, nil, false, nil)
	testTryCatch(a, func() error { return nil }, true, nil)
	testTryCatch(a, func() error { return errors.New("test") }, false, nil)
	testTryCatch(a, func() error { return nil }, true, nil, nil)
}

func testTryCatch(
	a *assert.Assertion,
	try func() error,
	isSuccess bool,
	catch func(error),
	finally ...func(),
) {
	isTryRun := false
	isCatchRun := false
	isHasFinally := false
	isFinallyRun := false

	finallyPart := make([]func(), 0, 1)
	if len(finally) > 0 {
		isHasFinally = true
		finallyPart = append(finallyPart, Conditional[func()](
			finally[0] != nil,
			func() {
				isFinallyRun = true
				finally[0]()
			},
			finally[0],
		))
	}

	err := TryCatch(
		Conditional(
			try != nil,
			func() error {
				isTryRun = true
				return try()
			},
			try,
		),
		Conditional[func(error)](
			catch != nil,
			func(err error) {
				isCatchRun = true
				catch(err)
			},
			catch,
		),
		finallyPart...,
	)

	if try != nil {
		a.TrueNow(isTryRun, "Expected try statement run")
	}

	if isSuccess {
		a.NotTrueNow(isCatchRun, "Unexpected catch statement run")
		a.NilNow(err)
	} else if catch != nil {
		a.TrueNow(isCatchRun, "Expected catch statement run")
		a.NotNilNow(err)
	}

	if isHasFinally {
		if finally[0] != nil {
			a.TrueNow(isFinallyRun, "Expected finally statement run")
		}
	} else {
		a.NotTrueNow(isFinallyRun, "Unexpected finally statement run")
	}
}
