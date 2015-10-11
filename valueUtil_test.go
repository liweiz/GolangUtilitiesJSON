package utilities

import (
	"fmt"
	"reflect"
	"testing"
)

type InOutFindTypeForValue struct {
	In  interface{}
	Out reflect.Kind
}

func TestFindTypeForValue(t *testing.T) {
	toTest := []InOutFindTypeForValue{
		// String
		InOutFindTypeForValue{"a", reflect.String},
		// Bool
		InOutFindTypeForValue{false, reflect.Bool},
		// Float64
		InOutFindTypeForValue{4.1, reflect.Float64},
		// Map
		InOutFindTypeForValue{map[string]interface{}{"a": 2}, reflect.Map},
		// Slice
		InOutFindTypeForValue{[]string{"b"}, reflect.Slice},
		// String reflect
		InOutFindTypeForValue{reflect.ValueOf("a"), reflect.String},
		// Bool reflect
		InOutFindTypeForValue{reflect.ValueOf(false), reflect.Bool},
		// Float64 reflect
		InOutFindTypeForValue{reflect.ValueOf(4.1), reflect.Float64},
		// Map reflect
		InOutFindTypeForValue{reflect.ValueOf(map[string]interface{}{"a": 2}), reflect.Map},
		// Slice reflect
		InOutFindTypeForValue{reflect.ValueOf([]string{"b"}), reflect.Slice},
	}
	for _, x := range toTest {
		k := FindTypeForValue(x.In)
		if k == x.Out {
			fmt.Println("TestFindTypeForValue passed: ", k)
		} else {
			t.Errorf("ERR: interface{}: %+v TestFindTypeForValue should be %+v. Not %+v\n", x.In, x.Out, k)
		}
	}
}

type InOutTestLowestReflectValue struct {
	In  interface{}
	Out reflect.Value
}

func TestLowestReflectValue(t *testing.T) {
	toTest := []InOutTestLowestReflectValue{
		// Non reflect.Value
		InOutTestLowestReflectValue{"a", reflect.ValueOf("a")},
		// 1 level reflect.Value
		InOutTestLowestReflectValue{reflect.ValueOf("a"), reflect.ValueOf("a")},
		// 5 level reflect.Value
		InOutTestLowestReflectValue{reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf(reflect.ValueOf("a"))))), reflect.ValueOf("a")},
	}
	for _, x := range toTest {
		v := LowestReflectValue(x.In)
		if v.String() == x.Out.String() {
			fmt.Println("TestLowestReflectValue passed: ", v)
		} else {
			t.Errorf("ERR: interface{}: %+v LowestReflectValue should be %+v. Not %+v\n", x.In, x.Out, v)
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
