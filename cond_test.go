package utils

import (
	"testing"

	"github.com/ghosind/go-assert"
)

func TestConditional(t *testing.T) {
	a := assert.New(t)
	language := "fr"

	ret := Conditional(language == "en", "Hello", "Bonjour")
	a.EqualNow(ret, "Bonjour")
}

func TestConditionalExpr(t *testing.T) {
	a := assert.New(t)

	var str *string

	ret := ConditionalExpr(
		str != nil,
		func() string {
			return *str
		},
		func() string {
			return ""
		},
	)
	a.Equal(ret, "")

	str = Pointer("Hello world")
	ret = ConditionalExpr(
		str != nil,
		func() string {
			return *str
		},
		func() string {
			return ""
		},
	)
	a.Equal(ret, "Hello world")
}

func TestMax(t *testing.T) {
	assertion := assert.New(t)

	a := 1
	b := 2

	max := Max(a, b)
	assertion.Equal(max, b)
}

func TestMin(t *testing.T) {
	assertion := assert.New(t)
	a := 1
	b := 2

	min := Min(a, b)
	assertion.Equal(min, a)
}
