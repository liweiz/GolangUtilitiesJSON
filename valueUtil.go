package utilities

import (
	"fmt"
	"reflect"
)

// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

// FindTypeForValue finds out the type and returns casted value in the form of reflect.Value for non collective data. It handles a data itself or in the form after one reflect.ValueOf() operation.
func FindTypeForValue(value interface{}) reflect.Kind {
	v := LowestReflectValue(value)
	switch v.Kind() {
	case reflect.Slice:
		if v.Len() > 0 {
			return reflect.Slice
		}
		return reflect.Invalid
	case reflect.Map:
		if len(v.MapKeys()) > 0 {
			return reflect.Map
		}
		return reflect.Invalid
	case reflect.String:
		return reflect.String
	case reflect.Float64:
		return reflect.Float64
	case reflect.Bool:
		return reflect.Bool
	default:
		fmt.Printf("ERROR: FindTypeForValue Value:%+v is out of current options. %+v\n", value, reflect.TypeOf(value))
		return reflect.Invalid
	}
}

// LowestReflectValue gets the closest reflect.Value for an interface{} value.
func LowestReflectValue(value interface{}) reflect.Value {
	v, ok := value.(reflect.Value)
	if ok {
		vv, cool := v.Interface().(reflect.Value)
		if cool {
			return LowestReflectValue(vv)
		}
		return v
	}
	return reflect.ValueOf(value)
}

// CompareValues compares two values. Currently targeted at value from map.
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
