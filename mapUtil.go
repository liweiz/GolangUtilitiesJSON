package utilities

import "reflect"

type StringKeyMap map[string]interface{}

type StringKeyMapType map[string]reflect.Kind

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
