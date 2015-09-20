package utilities

import (
	"fmt"
	"reflect"
)

// CompareValues compares two values. Currently targeted at value from map.
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null
func CompareValues(value1 interface{}, value2 interface{}) bool {
	if value1 == nil || value2 == nil {
		return false
	}
	switch value1 := value1.(type) {
	case map[string]interface{}:
		value2, ok := value2.(map[string]interface{})
		if ok {
			return CompareStringKeyMaps(value1, value2)
		}
	case bool:
		value2, ok := value2.(bool)
		if ok {
			if value1 == value2 {
				return true
			}
		}
	case float64:
		value2, ok := value2.(float64)
		if ok {
			if value1 == value2 {
				return true
			}
		}
	case string:
		value2, ok := value2.(string)
		if ok {
			if value1 == value2 {
				return true
			}
		}
	case []string:
		value2 := value2.([]string)
		return CompareStringSlices(value1, value2)
	case []map[string]interface{}:
		// In the context of JSON, interface{} must be one of the types of values from json.Unmarshal. We take care of map[string]interface{} specificly here.
		value2, ok := value2.([]map[string]interface{})
		if ok {
			return CompareStringKeyMapSlices(value1, value2)
		}
	default:
		fmt.Printf("ERROR: Type:%+v Value:%+v for key is out of current options.\n", reflect.TypeOf(value1), reflect.ValueOf(value1))
	}
	return false
}
