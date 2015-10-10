package utilities

import (
	"fmt"
	"reflect"
	"testing"
)

type InOutTestFindTypeForValue struct {
	InV     interface{}
	OutKind reflect.Kind
}

func TestFindTypeForValue(t *testing.T) {
	toTest := []InOutTestFindTypeForValue{
		// Float64
		InOutTestFindTypeForValue{float64(1), reflect.Float64},
		// String
		InOutTestFindTypeForValue{"a", reflect.String},
		// Bool
		InOutTestFindTypeForValue{false, reflect.Bool},
		// Map
		InOutTestFindTypeForValue{map[string]interface{}{"a": 1}, reflect.Map},
		// Slice
		InOutTestFindTypeForValue{[]interface{}{1}, reflect.Slice},
		// Nil
		InOutTestFindTypeForValue{nil, reflect.Invalid},
	}
	for _, x := range toTest {
		k := FindTypeForValue(x.InV)
		if k != x.OutKind {
			t.Errorf("ERR: value: %+v kind should be %+v. Not %+v\n", x.InV, x.OutKind, k)
		} else {
			fmt.Println("TestFindTypeForValue passed: ", k)
		}
	}
}

type InOutTestCompareValues struct {
	Ins []interface{}
	Out bool
}

func TestCompareValues(t *testing.T) {
	toTest := []InOutTestCompareValues{
		InOutTestCompareValues{
			[]interface{}{nil, m9},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{m9, m9},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{m9, m1},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{true, true},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{true, false},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{0.1, 0.1},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{0.1, 0.2},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{"m9", "m9"},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{"m9", ""},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{s1, s1},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{s1, s2},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1, m3}},
			true,
		},
		InOutTestCompareValues{
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1}},
			false,
		},
		InOutTestCompareValues{
			[]interface{}{[]float64{0.8, 0.9}, []float64{0.8, 0.9}},
			false,
		},
	}
	for _, x := range toTest {
		if CompareValues(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: interface{}: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		} else {
			fmt.Println("TestCompareValues passed: ", x.Ins[0], x.Ins[1])
		}
	}
}
