package utilities

import "reflect"

var InOutFindValueTypesInMap struct {
	Key         string
	V           reflect.Value
	noOfArrayLv int
	sec         StringKeyMapTypeSection
	isMap       bool
}

var aMap = map[string]interface{}{
	"a": 1,
	"b": "ss",
	"c": false,
	"d": nil,
	"e": []int{3},
	"f": map[string]int{"ff": 4},
}
