package utilities

import (
	"fmt"
	"reflect"
	"testing"
)

type InOutFindValueTypesInMap struct {
	InKey         string
	InV           reflect.Value
	InNoOfArrayLv int
	OutSec        StringKeyMapTypeSection
	OutIsMap      bool
}

var OneMap = map[string]interface{}{
	// Float64
	"a": float64(1),
	// String
	"b": "ss",
	// Bool
	"c": false,
	// Nil
	"d": nil,
	// Slice
	"e": []interface{}{3},
	// Map
	"f": map[string]interface{}{"ff": 4},
}

func TestFindValueTypesInUnknownMap(t *testing.T) {
	fmt.Println("\nSTART TestFindValueTypesInUnknownMap")
	toTest := []InOutFindValueTypesInMap{
		InOutFindValueTypesInMap{
			"anyKey",
			reflect.ValueOf(OneMap),
			999,
			StringKeyMapTypeSection{
				"anyKey",
				999,
				map[string]reflect.Kind{
					"a": reflect.Float64,
					"b": reflect.String,
					"c": reflect.Bool,
					"d": reflect.Invalid,
					"e": reflect.Slice,
					"f": reflect.Map,
				},
				reflect.Map,
			},
			true,
		},
	}
	for _, x := range toTest {
		s, m := FindValueTypesInUnknownMap(x.InKey, x.InV, x.InNoOfArrayLv)
		fmt.Println("TestFindValueTypesInUnknownMap start")
		if m == x.OutIsMap {
			fmt.Println("TestFindValueTypesInUnknownMap isMap passed")
			if CompareValues(s.TypeMap, x.OutSec.TypeMap) && s.ParentKey == x.OutSec.ParentKey && s.NoOfArrayLv == x.OutSec.NoOfArrayLv {
				fmt.Println("TestFindValueTypesInUnknownMap section passed")
			} else {
				t.Errorf("ERR: value: %+v should turn out to be section %+v. Not %+v\n", x.InV, x.OutSec, s)
			}
		} else {
			t.Errorf("ERR: value: %+v kind == map should be %+v. Not %+v\n", x.InV, x.OutIsMap, m)
		}
	}
}
