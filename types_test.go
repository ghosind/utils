package utils

import (
	"reflect"
	"testing"

	"github.com/ghosind/go-assert"
)

type testStruct1 struct {
	A int
}

type testStruct2[T any] struct {
	A T
}

func TestIsSameType(t *testing.T) {
	a := assert.New(t)

	t1 := new(testStruct1)
	t2 := testStruct1{}
	a.NotTrueNow(IsSameType(t1, t2))

	t3 := new(testStruct2[int])
	a.NotTrueNow(IsSameType(t1, t3))

	var t4 *testStruct1
	a.TrueNow(IsSameType(t1, t4))

	a.NotTrueNow(IsSameType(t1, nil))
}

func TestIsSameRawType(t *testing.T) {
	a := assert.New(t)

	t1 := new(testStruct1)
	t2 := testStruct1{}

	a.TrueNow(IsSameRawType(t1, t2))

	t3 := new(testStruct2[int])
	a.NotTrueNow(IsSameRawType(t1, t3))

	var t4 *testStruct1
	a.TrueNow(IsSameRawType(t1, t4))
}

func TestTypeOf(t *testing.T) {
	a := assert.New(t)

	ty := TypeOf(nil)
	a.EqualNow(ty, "<nil>")

	v1 := 0
	ty = TypeOf(v1)
	a.EqualNow(ty, "int")

	v2 := &v1
	ty = TypeOf(v2)
	a.EqualNow(ty, "*int")

	v3 := testStruct1{}
	ty = TypeOf(v3)
	a.EqualNow(ty, "utils.testStruct1")

	ty = TypeOf(&v3)
	a.EqualNow(ty, "*utils.testStruct1")
}

func TestRawTypeOf(t *testing.T) {
	a := assert.New(t)

	vType := "int"
	v1 := 0 // int value
	ty := RawTypeOf(v1)
	a.EqualNow(ty, vType)

	v2 := &v1 // pointer
	ty = RawTypeOf(v2)
	a.EqualNow(ty, vType)

	v3 := &v2 // pointer to pointer
	ty = RawTypeOf(v3)
	a.EqualNow(ty, vType)

	ty = RawTypeOf(nil)
	a.EqualNow(ty, "<nil>")

	v4 := &testStruct1{}
	ty = RawTypeOf(v4)
	a.EqualNow(ty, "utils.testStruct1")
}

func TestGetElem(t *testing.T) {
	a := assert.New(t)

	t1 := testStruct1{}
	t2 := &t1

	a.NotEqualNow(reflect.TypeOf(GetElem(t2)).Kind(), reflect.Ptr)

	t2 = nil
	v := GetElem(t2)
	a.NilNow(v)

	v = GetElem(nil)
	a.NilNow(v)

	t3 := 1
	v = GetElem(t3)
	a.Equal(v, t3)

	t4 := &t3
	v = GetElem(t4)
	a.Equal(v, t3)
}
