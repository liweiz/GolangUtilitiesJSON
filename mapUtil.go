package utilities

import (
	"fmt"
	"reflect"
)

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
						r[k.String] = ConvertMap(x)
					} else if t == reflect.Slice {
						r[k.String] = 
					}
				} else {
					switch t {
					case reflect.Bool:
						r[k.String] = x.Bool()
					case reflect.Float64:
						r[k.String] = x.Float()
					case reflect.String:
						r[k.String] = x.String()
					case reflect.Invalid:
						r[k.String] = nil
					default:
						fmt.Printf("ERROR: ConvertMap: out of current options.\n")
					}
				}

			}
		}
	}
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
