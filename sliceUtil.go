package utilities

import "reflect"

// ConvertToValueSlice converts reflect.Value to map[string]reflect.Value.
func ConvertToValueSlice(s reflect.Value) (valueSlice []reflect.Value, isSlice bool) {
	r := []reflect.Value{}
	if s.Kind() == reflect.Slice {
		if s.Len() > 0 {
			for i := 0; i < s.Len(); i++ {
				r = append(r, s.Index(i))
			}
			return r, true
		}
	}
	return r, false
}

// ConvertToBoolSlice returns true and []bool, when non-zero []bool is successfully converted.
func ConvertToBoolSlice(s []reflect.Value, typeChecked bool) (bool, []bool) {
	if len(s) > 0 {
		ok := false
		if typeChecked {
			ok = true
		} else {
			if reflect.TypeOf(s[0]).Kind() == reflect.Bool {
				ok = true
			}
		}
		if ok {
			r := []bool{}
			for _, x := range s {
				r = append(r, x.Bool())
			}
			return true, r
		}

	}
	return false, nil
}

// ConvertToStringSlice returns true and []string, when non-zero []string is successfully converted.
func ConvertToStringSlice(s []reflect.Value, typeChecked bool) (bool, []string) {
	if len(s) > 0 {
		ok := false
		if typeChecked {
			ok = true
		} else {
			if reflect.TypeOf(s[0]).Kind() == reflect.String {
				ok = true
			}
		}
		if ok {
			r := []string{}
			for _, x := range s {
				r = append(r, x.String())
			}
			return true, r
		}

	}
	return false, nil
}

// ConvertToFloat64Slice returns true and []float64, when non-zero []float64 is successfully converted.
func ConvertToFloat64Slice(s []reflect.Value, typeChecked bool) (bool, []float64) {
	if len(s) > 0 {
		ok := false
		if typeChecked {
			ok = true
		} else {
			if reflect.TypeOf(s[0]).Kind() == reflect.Float64 {
				ok = true
			}
		}
		if ok {
			r := []float64{}
			for _, x := range s {
				r = append(r, x.Float())
			}
			return true, r
		}

	}
	return false, nil
}

// ConvertToStringKeyMapSlice returns true and []map[string]interface, when non-zero []map[string]interface is successfully converted.
func ConvertToStringKeyMapSlice(s []reflect.Value, typeChecked bool) (bool, []map[string]reflect.Value) {
	if len(s) > 0 {
		ok := false
		if typeChecked {
			ok = true
		} else {
			if reflect.TypeOf(s[0]).Kind() == reflect.Map {
				if reflect.TypeOf(s[0]).Key().Kind() == reflect.String {
					ok = true
				}
			}
		}
		if ok {
			r := []map[string]reflect.Value{}
			for _, x := range s {
				r = append(r, ConvertToValueMap(x))
			}
			return true, r
		}
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

// CompareIntSlices finds out if two slices of strings are identical.
func CompareIntSlices(slice1 []int, slice2 []int) bool {
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
