package utils

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	v := 1

	vp := Pointer(v)
	if reflect.TypeOf(vp).Kind() != reflect.Ptr {
		t.Errorf("vp should be a pointer")
	}
	if v != *vp {
		t.Errorf("*vp(%d) should equals to v(%d)", *vp, v)
	}
}

func TestValue(t *testing.T) {
	v := 1
	vp := &v

	pv := Value(vp)
	if pv != v {
		t.Errorf("pv(%d) should equals to v(%d)", pv, v)
	}

	vp = nil
	pv = Value(vp)
	if pv != 0 {
		t.Errorf("pv should be 0, but %d", pv)
	}
}

func TestValueWithDefault(t *testing.T) {
	defaultVal := 2
	v := 1
	vp := &v

	pv := ValueWithDefault(vp, defaultVal)
	if pv != v {
		t.Errorf("pv(%d) should equals to v(%d)", pv, v)
	}

	vp = nil
	pv = ValueWithDefault(vp, defaultVal)
	if pv != defaultVal {
		t.Errorf("pv should be %d, but %d", defaultVal, pv)
	}
}

func TestPointerSlice(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6}

	ret := PointerSlice(arr)
	for i, p := range ret {
		if arr[i] != *p {
			t.Errorf("*ret[%d] is %d, expect %d", i, *p, arr[i])
		}
	}
}

func TestValueSlice(t *testing.T) {
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
			if *e != v {
				t.Errorf("ret[%d] is %d, expect %d", i, v, *e)
			}
		} else if v != zero {
			t.Errorf("ret[%d] is %d, expect %d", i, v, zero)
		}
	}
}

func TestPointerMap(t *testing.T) {
	greetings := map[string]string{
		"en":    "Hello",
		"fr":    "Bonjour",
		"zh_cn": "你好",
	}

	ret := PointerMap(greetings)
	for k, v := range greetings {
		if *ret[k] != v {
			t.Errorf("ret[%s] is %s, expect %s", k, *ret[k], v)
		}
	}
	for k := range ret {
		_, ok := greetings[k]
		if !ok {
			t.Errorf("Unexpected key %s in return value", k)
		}
	}
}

func TestValueMap(t *testing.T) {
	greetings := map[string]*string{
		"en":    Pointer("Hello"),
		"fr":    nil,
		"zh_cn": Pointer("你好"),
	}

	ret := ValueMap(greetings)
	for k, v := range greetings {
		if v == nil {
			_, ok := ret[k]
			if ok {
				t.Errorf("Unexpected key %s in return value", k)
			}
		} else if *v != ret[k] {
			t.Errorf("ret[%s] is %s, expect %s", k, ret[k], *v)
		}
	}
	for k := range ret {
		_, ok := greetings[k]
		if !ok {
			t.Errorf("Unexpected key %s in return value", k)
		}
	}
}
