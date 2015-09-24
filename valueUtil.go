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






func ExploreTypeForValue(value interface{}) (reflect.Kind, reflect.Value){
	v := reflect.ValueOf(value)
	if v.IsNil {
		return reflect.Invalid, nil
	} else {
		switch v.Kind() {
		case reflect.Slice:
			return reflect.Slice, v
		case reflect.Map:
			return reflect.Map, v
		case reflect.String:
			return reflect.String, v
		case reflect.Float64:
			return reflect.Float64, v
		case reflect.Bool:
			return reflect.Bool, v
		default:
			return reflect.Invalid, nil
		}
	}
}

func DispatchToValueExplorer(valueType reflect.Kind, value reflect.Value) {
	switch valueType {
	case reflect.Slice:
		ok, o := ExploreSliceValue(value)
		if ok {
			switch o.Type().Elem() {
			case reflect.Bool:
				ok, r := ConvertToBoolSlice(o)
				if ok {

				}
			}
		}
		return reflect.Slice, v
	case reflect.Map:
		return reflect.Map, v
	case reflect.String:
		return reflect.String, v
	case reflect.Float64:
		return reflect.Float64, v
	case reflect.Bool:
		return reflect.Bool, v
	default:
		return reflect.Invalid, nil
	}
}




func ExploreSliceValue(value reflect.Value) (isValid bool, output []reflect.Value) {
	r := []reflect.Value{}
	t := value.Type()
	if t.Kind() == reflect.Slice {
		l := t.Len()
		if l > 0 {
			v := reflect.ValueOf(value)
			for i := 0; i < l; i++ {
				append(r, v.Index(i))
			}
			return true, r
		}
	}
	return false, nil
}

func ExploreStringMapValue(value interface{}) (isValid bool, output map[string]reflect.Value) {
	r := map[string]reflect.Value{}
	t := value.Type()
	if t.Kind() == reflect.Map {
		l := len(t)
		if l > 0 {
			v := reflect.ValueOf(value)
			k := v.MapKeys()
			if k[0].Kind() == reflect.String {
				for _, x := range k {
					r[x.String] = v.MapIndex(x)
				}
				return true, r
			}
		}
	}
	return false, nil
}



// FindTypeForValue returns type for the value. If it's a map, further action will be needed. If it's a nil, we leave it to the next check.
func FindTypeForValue(value interface{}) (valueType reflect.Kind, furtherActionNeeded bool) {
	if value == nil {
		return reflect.Invalid, false
	}
	switch reflect.TypeOf(value).Kind() {
	case map[string]interface{}:
		return reflect.Map, true
	case bool:
		return reflect.Bool, false
	case float64:
		return reflect.Float64, false
	case string:
		return reflect.String, false
	default:
		if reflect.TypeOf(value).len() > 0 {
			return reflect.Slice, true
		}
		fmt.Printf("ERROR: Type:%+v Value:%+v for key is out of current options.\n", reflect.TypeOf(value1), reflect.ValueOf(value1))
		return reflect.Invalid, false
	}
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
