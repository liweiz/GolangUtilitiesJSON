package utilities

import "reflect"

// ConvertToValueMap converts reflect.Value to map[string]reflect.Value.
func ConvertToValueMap(m reflect.Value) map[string]reflect.Value {
	r := map[string]reflect.Value{}
	if m.Kind() == reflect.Map {
		if m.Len() > 0 {
			keys := m.MapKeys()
			for _, k := range keys {
				r[k.String()] = m.MapIndex(k)
			}
		}
	}
	return r
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
