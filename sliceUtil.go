package utilities

import "reflect"

// ConvertToBoolSlice returns true and []bool, when non-zero []bool is successfully converted.
func ConvertToBoolSlice(s []reflect.Value) (bool, []bool) {
	if len(s) > 0 {
		if reflect.TypeOf(s[0]).Kind() == reflect.Bool {
			r := []bool{}
			for _, x := range s {
				// append(r, x.Bool())
			}
			return true, r
		}
	}
	return false, nil
}

// ConvertToStringSlice returns true and []string, when non-zero []string is successfully converted.
func ConvertToStringSlice(s []reflect.Value) (bool, []string) {
	if len(s) > 0 {
		if reflect.TypeOf(s[0]).Kind() == reflect.Bool {
			r := []string{}
			for _, x := range s {
				// append(r, x.String())
			}
			return true, r
		}
	}
	return false, nil
}

// ConvertToFloat64Slice returns true and []float64, when non-zero []float64 is successfully converted.
func ConvertToFloat64Slice(s []reflect.Value) (bool, []float64) {
	if len(s) > 0 {
		if reflect.TypeOf(s[0]).Kind() == reflect.Float64 {
			r := []float64{}
			for _, x := range s {
				// append(r, x.Float())
			}
			return true, r
		}
	}
	return false, nil
}

// ConvertToStringKeyMapSlice returns true and []map[string]interface, when non-zero []map[string]interface is successfully converted.
func ConvertToStringKeyMapSlice(s []reflect.Value) (bool, []map[string]interface{}) {
	if len(s) > 0 {
		if reflect.TypeOf(s[0]).Kind() == reflect.Map {
			r := []float64{}
			for _, x := range s {
				// append(r, x.Float())
			}
			// return true, r
		}
	}
	return false, nil
}

// CheckIfIsSliceValue identifies if the value is a slice.
func CheckIfIsSliceValue(value interface{}) (bool, []interface{}) {
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Slice {
		// if t.Elem()
	}
	return false, nil
}

// CompareStringKeyMapSlices finds out if two slices are identical.
func CompareStringKeyMapSlices(slice1 []map[string]interface{}, slice2 []map[string]interface{}) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if !CompareValues(x, slice2[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// CheckIfIsStringKeyMapSlice finds out if it's []map[string]interface{} and return a []map[string]interface{} casted from a.
func CheckIfIsStringKeyMapSlice(s []interface{}) (bool, []map[string]interface{}) {
	if len(s) == 0 {
		return false, nil
	}
	newS := []map[string]interface{}{}
	for i, x := range s {
		_, ok := x.(map[string]interface{})
		if !ok {
			return false, nil
		}
		y, good := s[i].(map[string]interface{})
		if !good {
			return false, nil
		}
		newS = append(newS, y)
	}
	return true, newS
}

// CompareStringSlices finds out if two slices of strings are identical.
func CompareStringSlices(slice1 []string, slice2 []string) bool {
	if len(slice1) == len(slice2) {
		for i, x := range slice1 {
			if x != slice2[i] {
				return false
			}
		}
		return true
	}
	return false
}
