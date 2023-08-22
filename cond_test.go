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

func TestMaxN(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(MaxN[int](), 0)
	assertion.Equal(MaxN(2, 1, 1), 2)
	assertion.Equal(MaxN(-1, 1, -1), 1)
	assertion.Equal(MaxN(1, 2, 3, 4, 3, 2, 1), 4)
}

func TestMin(t *testing.T) {
	assertion := assert.New(t)
	a := 1
	b := 2

	min := Min(a, b)
	assertion.Equal(min, a)
}

func TestMinN(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(MinN[int](), 0)
	assertion.Equal(MinN(2, 1, 2), 1)
	assertion.Equal(MinN(-1, 1, -1), -1)
	assertion.Equal(MinN(4, 3, 2, 1, 2, 3, 4), 1)
}
