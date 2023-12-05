package utils

import (
	"reflect"
	"testing"

	"github.com/ghosind/go-assert"
)

func TestPointer(t *testing.T) {
	a := assert.New(t)

	v := 1

	vp := Pointer(v)
	a.EqualNow(reflect.TypeOf(vp).Kind(), reflect.Ptr)
	a.NotNilNow(vp)
	a.EqualNow(*vp, v)
}

func TestPointerWithDefault(t *testing.T) {
	a := assert.New(t)

	v := Pointer(1)
	dv := Pointer(2) // default value pointer
	a.EqualNow(PointerWithDefault(v, dv), v)

	v = nil
	a.EqualNow(PointerWithDefault(v, dv), dv)

	dv = nil
	a.NilNow(PointerWithDefault(v, dv))
}

func TestValue(t *testing.T) {
	a := assert.New(t)

	v := 1
	vp := &v

	pv := Value(vp)
	a.EqualNow(pv, v)

	vp = nil
	pv = Value(vp)
	a.EqualNow(pv, 0)
}

func TestValueWithDefault(t *testing.T) {
	a := assert.New(t)

	defaultVal := 2
	v := 1
	vp := &v

	pv := ValueWithDefault(vp, defaultVal)
	a.EqualNow(pv, v)

	vp = nil
	pv = ValueWithDefault(vp, defaultVal)
	a.EqualNow(pv, defaultVal)
}

func TestPointerSlice(t *testing.T) {
	a := assert.New(t)

	arr := []int{0, 1, 2, 3, 4, 5, 6}

	ret := PointerSlice(arr)
	for i, p := range ret {
		a.NotNilNow(p)
		a.EqualNow(*p, arr[i])
	}
}

func TestValueSlice(t *testing.T) {
	a := assert.New(t)

	arr := []*int{
		Pointer(1),
		Pointer(2),
		Pointer(3),
		nil,
		Pointer(5),
		Pointer(6),
	}

	ret := ValueSlice(arr)
	zero := 0
	for i, e := range arr {
		v := ret[i]
		if e != nil {
			a.EqualNow(*e, v)
		} else {
			a.EqualNow(v, zero)
		}
	}
}

func TestPointerMap(t *testing.T) {
	a := assert.New(t)

	greetings := map[string]string{
		"en":    "Hello",
		"fr":    "Bonjour",
		"zh_cn": "你好",
	}

	ret := PointerMap(greetings)
	a.EqualNow(len(ret), len(greetings))
	for k, v := range greetings {
		a.NotNilNow(ret[k])
		a.EqualNow(*ret[k], v)
	}
	for k := range ret {
		_, ok := greetings[k]
		a.TrueNow(ok)
	}
}

func TestValueMap(t *testing.T) {
	a := assert.New(t)

	greetings := map[string]*string{
		"en":    Pointer("Hello"),
		"fr":    nil,
		"zh_cn": Pointer("你好"),
	}

	ret := ValueMap(greetings)
	for k, v := range greetings {
		if v == nil {
			_, ok := ret[k]
			a.NotTrueNow(ok)
		} else if *v != ret[k] {
			a.EqualNow(ret[k], *v)
		}
	}
	for k := range ret {
		_, ok := greetings[k]
		a.TrueNow(ok)
	}
}
