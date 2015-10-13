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

// Data type process flow:
// 1. loaded in as interface{}
// 2. find out type
// 2.1 primitive type
// 2.2 map/slice of interface{} or nil

// FindTypeForValue finds out the type. It handles a data itself and the data after any times of operations in reflect.ValueOf().
func FindTypeForValue(value interface{}) reflect.Kind {
	v := LowestReflectValue(value)
	fmt.Println("FindTypeForValue tttttt: ", v.Kind().String())
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
	case reflect.Interface:
		fmt.Println("FindTypeForValue Interface ENTRY")
		switch v.Interface().(type) {
		case bool:
			fmt.Println("FindTypeForValue Interface: Bool")
			return reflect.Bool
		case string:
			fmt.Println("FindTypeForValue Interface: String")
			return reflect.String
		case float64:
			fmt.Println("FindTypeForValue Interface: Float64")
			return reflect.Float64
		case map[string]interface{}:
			fmt.Println("FindTypeForValue Interface: map[string]int")
			return reflect.Map
		case []interface{}:
			fmt.Println("FindTypeForValue Interface: []int")
			return reflect.Slice
		default:
			fmt.Println("FindTypeForValue Interface: Invalid")
			return reflect.Invalid
		}
	default:
		fmt.Printf("ERROR: FindTypeForValue Value:%+v is out of current options. %+v\n", value, reflect.TypeOf(value))
		return reflect.Invalid
	}
}

// func InterfaceIsMap(i interface{}) bool {
//
// }

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

// CompareValues compares two values. Currently targeted at value in bool/string/float64/reflect.Kind and complex value map/slice derived from them.
func CompareValues(value1 interface{}, value2 interface{}) bool {
	if value1 == nil || value2 == nil {
		return false
	}
	switch valueA := value1.(type) {
	case map[string]reflect.Kind:
		valueB, ok := value2.(map[string]reflect.Kind)
		if ok {
			var value1A, value2B map[string]interface{}
			for s1, x1 := range valueA {
				value1A[s1] = interface{}(x1)
			}
			for s2, x2 := range valueB {
				value2B[s2] = interface{}(x2)
			}
			return CompareStringKeyMaps(value1A, value2B)
		}
	case map[string]interface{}:
		valueB, ok := value2.(map[string]interface{})
		if ok {
			return CompareStringKeyMaps(valueA, valueB)
		}
	case reflect.Kind:
		valueB, ok := value2.(reflect.Kind)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case bool:
		valueB, ok := value2.(bool)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case float64:
		valueB, ok := value2.(float64)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case string:
		valueB, ok := value2.(string)
		if ok {
			if valueA == valueB {
				return true
			}
		}
	case []string:
		valueB := value2.([]string)
		return CompareStringSlices(valueA, valueB)
	case []map[string]interface{}:
		// In the context of JSON, interface{} must be one of the types of values from json.Unmarshal. We take care of map[string]interface{} specificly here.
		valueB, ok := value2.([]map[string]interface{})
		if ok {
			return CompareStringKeyMapSlices(valueA, valueB)
		}
	default:
		fmt.Printf("ERROR CompareValues: Type:%+v Value:%+v for key is out of current options.\n", reflect.TypeOf(value1), reflect.ValueOf(value1))
	}
	return false
}
