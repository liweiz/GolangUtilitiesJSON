package utilities

// type InOutFindValueTypesInMap struct {
// 	InKey         string
// 	InV           reflect.Value
// 	InNoOfArrayLv int
// 	OutSec        StringKeyMapTypeSection
// 	OutIsMap      bool
// }
//
// var OneMap = map[string]interface{}{
// 	// Float64
// 	"a": float64(1),
// 	// String
// 	"b": "ss",
// 	// Bool
// 	"c": false,
// 	// Nil
// 	"d": nil,
// 	// Slice
// 	"e": []int{3},
// 	// Map
// 	"f": map[string]int{"ff": 4},
// }
//
// func TestFindValueTypesInUnknownMap(t *testing.T) {
// 	toTest := []InOutFindValueTypesInMap{
// 		InOutFindValueTypesInMap{
// 			"anyKey",
// 			reflect.ValueOf(OneMap),
// 			999,
// 			StringKeyMapTypeSection{
// 				"anyKey",
// 				999,
// 				map[string]reflect.Kind{
// 					"a": reflect.Float64,
// 					"b": reflect.String,
// 					"c": reflect.Bool,
// 					"d": reflect.Invalid,
// 					"e": reflect.Slice,
// 					"f": reflect.Map,
// 				},
// 				reflect.Map,
// 			},
// 			true,
// 		},
// 	}
// 	for _, x := range toTest {
// 		s, m := FindValueTypesInUnknownMap(x.InKey, x.InV, x.InNoOfArrayLv)
// 		fmt.Println("TestFindTypeForValue xxxxxx")
// 		if m == x.OutIsMap {
// 			fmt.Println("TestFindTypeForValue isMap passed")
// 			if CompareValues(s.TypeMap, x.OutSec.TypeMap) && s.ParentKey == x.OutSec.ParentKey && s.NoOfArrayLv == x.OutSec.NoOfArrayLv {
// 				fmt.Println("TestFindTypeForValue section passed")
// 			} else {
// 				t.Errorf("ERR: value: %+v should turn out to be section %+v. Not %+v\n", x.InV, x.OutSec, s)
// 			}
// 		} else {
// 			t.Errorf("ERR: value: %+v kind == map should be %+v. Not %+v\n", x.InV, x.OutIsMap, m)
// 		}
// 	}
// }
