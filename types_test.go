package utils

import (
	"reflect"
	"testing"
)

type testStruct1 struct {
	a int
}

type testStruct2[T any] struct {
	a T
}

func TestIsSameType(t *testing.T) {
	t1 := new(testStruct1)
	t2 := testStruct1{}
	if IsSameType(t1, t2) {
		t.Error("IsSameType returns true, expect false")
	}

	t3 := new(testStruct2[int])
	if IsSameType(t1, t3) {
		t.Error("IsSameType returns true, expect false")
	}

	var t4 *testStruct1
	if !IsSameType(t1, t4) {
		t.Error("IsSameType returns false, expect true")
	}
}

func TestIsSameRawType(t *testing.T) {
	t1 := new(testStruct1)
	t2 := testStruct1{}

	if !IsSameRawType(t1, t2) {
		t.Error("IsSameFawType returns false, expect true")
	}

	t3 := new(testStruct2[int])
	if IsSameRawType(t1, t3) {
		t.Error("IsSameRawType returns true, expect false")
	}

	var t4 *testStruct1
	if !IsSameRawType(t1, t4) {
		t.Error("IsSameRawType returns false, expect true")
	}
}

func TestGetElem(t *testing.T) {
	t1 := testStruct1{}
	t2 := &t1

	if reflect.TypeOf(GetElem(t2)).Kind() == reflect.Ptr {
		t.Error("GetElem(t2) returns a pointer, expect not pointer")
	}

	t2 = nil
	if v := GetElem(t2); v != nil {
		t.Errorf("GetElem(t2) returns %v, expect nil", v)
	}

	if v := GetElem(nil); v != nil {
		t.Errorf("GetElem(nil) returns %v, expect nil", v)
	}
}
