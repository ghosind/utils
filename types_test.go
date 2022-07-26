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
		t.Error("IsSameType(t1, t2) returns true, expect false")
	}

	t3 := new(testStruct2[int])
	if IsSameType(t1, t3) {
		t.Error("IsSameType(t1, t3) returns true, expect false")
	}

	var t4 *testStruct1
	if !IsSameType(t1, t4) {
		t.Error("IsSameType(t1, t4) returns false, expect true")
	}

	if IsSameType(t1, nil) {
		t.Error("IsSameType(t1, nil) returns true, expect false")
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

func TestTypeOf(t *testing.T) {
	if ty := TypeOf(nil); ty != "<nil>" {
		t.Errorf("TypeOf(nil) returns %s, expect <nil>", ty)
	}

	v1 := 0
	if ty := TypeOf(v1); ty != "int" {
		t.Errorf("TypeOf(v1) returns %s, expect int", ty)
	}

	v2 := &v1
	if ty := TypeOf(v2); ty != "*int" {
		t.Errorf("TypeOf(v2) returns %s, expect *int", ty)
	}
}

func TestRawTypeOf(t *testing.T) {
	vType := "int"
	v1 := 0 // int value
	if ty := RawTypeOf(v1); ty != vType {
		t.Errorf("RawTypeOf(v1) returns %s, expect %s", ty, vType)
	}
	v2 := &v1 // pointer
	if ty := RawTypeOf(v2); ty != vType {
		t.Errorf("RawTypeOf(v2) returns %s, expect %s", ty, vType)
	}
	v3 := &v2 // pointer to pointer
	if ty := RawTypeOf(v3); ty != vType {
		t.Errorf("RawTypeOf(v3) returns %s, expect %s", ty, vType)
	}

	if ty := RawTypeOf(nil); ty != "<nil>" {
		t.Errorf("RawTypeOf(v1) returns %s, expect <nil>", ty)
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

	t3 := 1
	if v := GetElem(t3); v != t3 {
		t.Errorf("GetElem(t3) returns %v, expect %d", v, t3)
	}

	t4 := &t3
	if v := GetElem(t4); v != t3 {
		t.Errorf("GetElem(t4) returns %v, expect %d", v, t3)
	}
}
