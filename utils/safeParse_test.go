package utils

import (
	"testing"
)

type Test struct {
	in  interface{}
	out interface{}
}

var testsInt = []Test{
	{int(1), 1},
	{int64(1), 1},
	{int32(1), 1},
	{int16(1), 1},
	{int8(1), 1},

	{uint(1), 1},
	{uint64(1), 1},
	{uint32(1), 1},
	{uint16(1), 1},
	{uint8(1), 1},

	{float64(1), 1},
	{float32(1), 1},
	{"1", 1},
	{true, 1},
	{false, 0},
	{nil, 0},
}

// Test Safe Parse Int
func TestSafeParseInt(t *testing.T) {
	for i, test := range testsInt {
		out, err := SafeParseInt(test.in)
		if err != nil {
			t.Errorf("#%d: Result %d Error : %s", i, out ,err)
		}
	}
}

var testsInt64 = []Test{
	{int(1), int64(1)},
	{int64(1), int64(1)},
	{int32(1), int64(1)},
	{int16(1), int64(1)},
	{int8(1), int64(1)},

	{uint(1), int64(1)},
	{uint64(1), int64(1)},
	{uint32(1), int64(1)},
	{uint16(1), int64(1)},
	{uint8(1), int64(1)},

	{float64(1), int64(1)},
	{float32(1), int64(1)},
	{"1", 1},
	{true, 1},
	{false, 0},
	{nil, 0},
}

// Test Safe Parse Int64
func TestSafeParseInt64(t *testing.T) {
	for i, test := range testsInt {
		out, err := SafeParseInt64(test.in)
		if err != nil {
			t.Errorf("#%d: Result %d Error : %s", i, out ,err)
		}
	}
}

var testsIndirect = []Test{
	{nil, nil},
	{1, 1},
}

func TestIndirect(t *testing.T) {
	for i, test := range testsIndirect {
		out :=  indirect(test.in)
		if out != test.out {
			t.Errorf("#%d: indirect(%#v)=%#v; want %#v", i, test.in, out, test.out)
		}
	}
}
