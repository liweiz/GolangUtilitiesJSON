package utilities

import "reflect"

// There are two dimensions in a map: 1) depth, a path to a value node 2) width, a section in a branch for a given level.

// StringKeyMap is the map counterpart for JSON.
type StringKeyMap map[string]interface{}

// StringKeyMapTypeMap stores StringKeyMap's value type for one layer of one branch.
type StringKeyMapTypeMap map[string]reflect.Kind

// StringKeyMapTypeSection stores the parent key and possible array index to link the StringKeyMapType info for its parent layer. An empty ParentKey indicates it's a root map. Index == -1 indicates parent level is not an array.
type StringKeyMapTypeSection struct {
	ParentKey string
	Index     int16
	TypeMap   StringKeyMapTypeMap
	TypeSlice reflect.Kind
}

// StringKeyMapTypeSections stores all the StringKeyMapTypeSections for a map schema.
var StringKeyMapTypeSections = []StringKeyMapTypeSection{}

// ScanOneBranchLevelMap finds out all value types in this level.
func ScanOneBranchLevelMap(key string, v reflect.Value) (isMap bool, section StringKeyMapTypeSection) {
	if v.Kind() == reflect.Map {
		r := StringKeyMapTypeSection{key, -1, StringKeyMapTypeMap{}, reflect.Invalid}
		for _, x := range v.MapKeys() {
			vv := v.MapIndex(x)
			known, t := CheckIfValueTypeAlreadyKnown(vv.String(), -1)
			if known {
				r.TypeMap[x.String()] = t
			} else {
				k, va := ExploreTypeForValue(vv)
				if k == reflect.Invalid {
					return false, StringKeyMapTypeSection{}
				}
				r.TypeMap[x.String()] = k
			}
		}
		return true, r
	}
	return false, StringKeyMapTypeSection{}
}

// CheckIfValueTypeAlreadyKnown finds out if the value's type is known. index is used to locate the value in an array.
func CheckIfValueTypeAlreadyKnown(key string, index ...int16) (known bool, valueType reflect.Kind) {
	var isInArray bool
	if len(index) == 0 {
		isInArray = false
	} else if index[0] < 0 {
		isInArray = false
	} else {
		isInArray = true
	}
	for _, x := range StringKeyMapTypeSections {
		if key == x.ParentKey {
			if isInArray {
				if index[0] == x.Index {
					return true, x.TypeSlice
				}
			} else {
				return true, x.TypeMap[key]
			}
		}
	}
	return false, reflect.Invalid
}

// ValuePath is the container for a level-based path to a value in a StringKeyMap. It also contains single and slice (if applicable) forms of value(s).
// Level here is not just identified by keys in different levels. It's also identified by slices.
// Keys is a slice to record the keys for all levels on the path. If the value is in an array, the corresponding key for that level is empty string: "" since empty string key in JSON seems have no real value. We can simply ignore it here.
// Indexes is a slice to record the index for all levels on the path. If the value is in a map, the corresponding index for that level is -1.
// ReflectValue is the corresponding value stored as a single value.
// ReflectValues is the corresponding value stored as a slice of values, which only applies to array in .
type ValuePath struct {
	Keys          []string
	Indexes       []int16
	ReflectValue  reflect.Value
	ReflectValues []reflect.Value
}

// ValueKey returns the key for the value.
func (p ValuePath) ValueKey() string {
	l := len(p.Keys)
	if l > 1 {
		if len(p.Keys[l-1]) == 0 {
			return p.Keys[l-2]
		}
		return p.Keys[l-1]
	}
	return p.Keys[0]
}

// ValueIndex returns the index for the value, if the value is in an array. Otherwise, it returns -1 to indicate the value is not from an array.
func (p ValuePath) ValueIndex() int16 {
	l := len(p.Keys)
	if l > 1 {
		if len(p.Keys[l-1]) == 0 {
			return p.Indexes[l-1]
		}
	}
	return -1
}

// GetNextLevelValue
func GetNextLevelValue(reflect.Value) reflect.Value {

}

func (p ValuePath) getReflectValueForPath(s StringKeyMap) {

}

func (s StringKeyMapType) GetValuePath(parentPath ValuePath, sMap StringKeyMap) ValuePath {
	v := s
	for _, k := range parentPath.Keys {
		if len(k) == 0 {

		} else {

		}
	}
}

func ConvertMap(m interface{}) map[string]interface{} {
	r := map[string]interface{}{}
	if reflect.TypeOf(m).Kind() == reflect.Map {
		if reflect.TypeOf(m).Len() > 0 {
			v := reflect.ValueOf(m)
			keys := v.MapKeys()
			for _, k := range keys {
				x := v.MapIndex(k)
				t, furtherActionNeeded := FindTypeForValue(x)
				if furtherActionNeeded {
					if t == reflect.Map {
						// r[k.String] = ConvertMap(x)
					} else if t == reflect.Slice {
						// r[k.String] =
					}
				} else {
					// switch t {
					// case reflect.Bool:
					// 	r[k.String] = x.Bool()
					// case reflect.Float64:
					// 	r[k.String] = x.Float()
					// case reflect.String:
					// 	r[k.String] = x.String()
					// case reflect.Invalid:
					// 	r[k.String] = nil
					// default:
					// 	fmt.Printf("ERROR: ConvertMap: out of current options.\n")
					// }
				}

			}
		}
	}
	return map[string]interface{}{}
}

// GetKeyMap explores all types for a map's keys. It's value is either reflect.Kind or a map[string]interface{}.
func GetKeyMap(jsonMap map[string]interface{}, typeMap map[string]interface{}) (bool, map[string]interface{}) {
	m := map[string]interface{}{}
	for key, x := range jsonMap {
		if typeMap[key] != nil && typeMap[key] != reflect.Invalid {
			m[key] = typeMap[key]
		} else {

		}
	}
	return false, map[string]interface{}{}
}

// CompareStringKeyMaps finds out if two maps with key typed in string are identical.
func CompareStringKeyMaps(map1 map[string]interface{}, map2 map[string]interface{}) bool {
	if len(map1) == len(map2) {
		for key1, value1 := range map1 {
			if !CompareValues(value1, map2[key1]) {
				return false
			}
		}
		return true
	}
	return false
}
