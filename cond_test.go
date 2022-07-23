package utils

import "testing"

func TestConditional(t *testing.T) {
	language := "fr"

	ret := Conditional(language == "en", "Hello", "Bonjour")
	if ret == "Hello" {
		t.Errorf("Conditional returns %s, expect \"Bonjour\"", ret)
	}
}

func TestMax(t *testing.T) {
	a := 1
	b := 2

	max := Max(a, b)
	if max == a {
		t.Errorf("Max returns %d, expect %d", max, b)
	}
}

func TestMin(t *testing.T) {
	a := 1
	b := 2

	min := Min(a, b)
	if min == b {
		t.Errorf("Min returns %d, expect %d", min, a)
	}
}
