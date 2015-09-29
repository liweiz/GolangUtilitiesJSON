package utilities

import "reflect"

// There are two dimensions in a map: 1) depth, a path to a value node 2) width, a section in a branch for a given level.

// StringKeyMap is the map counterpart for JSON.
type StringKeyMap map[string]interface{}

// StringKeyMapTypeMap stores StringKeyMap's value type for one layer of one branch.
type StringKeyMapTypeMap map[string]reflect.Kind

// StringKeyMapTypeSection stores the parent key and possible array index to link the StringKeyMapType info for its parent layer. An empty ParentKey indicates it's a root map. Index == -1 indicates parent level is not an array.
type StringKeyMapTypeSection struct {
	ParentKey string
	Indices   []int
	TypeMap   StringKeyMapTypeMap
	TypeSlice reflect.Kind
}

// StringKeyMapTypeSections stores all the StringKeyMapTypeSections for a map schema.
type StringKeyMapTypeSections []StringKeyMapTypeSection

var mapTree = StringKeyMapTypeSections{}

// FindElemInSections returns true with a valid value if a value is found.
func (s StringKeyMapTypeSections) FindElemInSections(key string, indices []int) (found bool, value StringKeyMapTypeSection) {
	for _, x := range s {
		if x.ParentKey == key && CompareIntSlices(x.Indices, indices) {
			return true, x
		}
	}
	return false, StringKeyMapTypeSection{}
}

// UpdateElemToSections updates/adds a StringKeyMapTypeSection in StringKeyMapTypeSections.
func (s StringKeyMapTypeSections) UpdateElemToSections(updated StringKeyMapTypeSection) {
	s.RemoveElemInSections(updated.ParentKey, updated.Indices)
	s = append(s, updated)
}

// RemoveElemInSections removes a StringKeyMapTypeSection from StringKeyMapTypeSections and assigns the new slice to StringKeyMapTypeSections.
func (s StringKeyMapTypeSections) RemoveElemInSections(key string, indices []int) {
	y := []StringKeyMapTypeSection{}
	for i, x := range s {
		if !(x.ParentKey == key && CompareIntSlices(x.Indices, indices)) {
			y = append(y, x)
		}
	}
	s = y
}

// func ScanStringKeyMap(m map[string]interface{}) (isMap bool, sections StringKeyMapTypeSections) {
// 	r := []StringKeyMapTypeSection{}
// 	if len(m) > 0 {
// 		ok, s := ScanOneBranchLevelMap("UtilityStringKeyMapRoot", reflect.ValueOf(m))
// 		if ok {
// 			for
//
// 		}
// 	}
//
// }

// FindValueTypesInSlice finds out all value types in this level in a slice.
func FindValueTypesInSlice(key string, v reflect.Value, indices []int) (isSlice bool, section StringKeyMapTypeSection) {
	if v.Kind() == reflect.Slice {
		if v.Len() > 0 {
			_, s := ConvertToValueSlice(v)
			k, _ := FindTypeForValue(s[0])
			return true, StringKeyMapTypeSection{key, indices, StringKeyMapTypeMap{}, k}
		}
	}
	return false, StringKeyMapTypeSection{}
}

// FindValueTypesInMap finds out all value types in this level in a map.
func FindValueTypesInMap(key string, v reflect.Value, indices []int) (isMap bool, section StringKeyMapTypeSection) {
	if v.Kind() == reflect.Map {
		i := reflect.Invalid
		if len(v.MapKeys()) > 0 {
			i = reflect.Map
		}
		r := StringKeyMapTypeSection{key, indices, StringKeyMapTypeMap{}, i}
		newTypeFound := false
		for _, x := range v.MapKeys() {
			known, t := CheckIfValueTypeAlreadyKnown(x.String(), indices)
			if known {
				r.TypeMap[x.String()] = t
			} else {
				k, _ := FindTypeForValue(v.MapIndex(x))
				if k == reflect.Invalid {
					return false, StringKeyMapTypeSection{}
				}
				r.TypeMap[x.String()] = k
				newTypeFound = true
			}
		}
		if newTypeFound {
			mapTree.RemoveElemInSections(key, indices)
		}
		return true, r
	}
	return false, StringKeyMapTypeSection{}
}

// CheckIfValueTypeAlreadyKnown finds out if the value's type is known. index is used to locate the value in an array.
func CheckIfValueTypeAlreadyKnown(key string, indices []int) (known bool, valueType reflect.Kind) {
	var isInArray bool
	if len(indices) == 0 {
		isInArray = false
	} else {
		isInArray = true
	}
	for _, x := range mapTree {
		if key == x.ParentKey {
			if isInArray {
				if CompareIntSlices(indices, x.Indices) {
					return true, x.TypeSlice
				}
			} else {
				return true, x.TypeMap[key]
			}
		}
	}
	return false, reflect.Invalid
}
