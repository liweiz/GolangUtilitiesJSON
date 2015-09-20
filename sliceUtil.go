package utilities

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
