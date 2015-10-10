package utilities

import (
	"fmt"
	"reflect"
)

// There are two dimensions in a map: 1) depth, a path to a value node 2) width, a section in a branch for a given level.

// StringKeyMapTypeSections stores all the StringKeyMapTypeSections for a map schema.
type StringKeyMapTypeSections []StringKeyMapTypeSection

// FindElemInSections returns true with a valid value if a value is found.
func (s StringKeyMapTypeSections) FindElemInSections(key string, noOfArrayLv int) (value StringKeyMapTypeSection, found bool) {
	for _, x := range s {
		if x.ParentKey == key && x.NoOfArrayLv == noOfArrayLv {
			return x, true
		}
	}
	return StringKeyMapTypeSection{}, false
}

// UpdateElemToSections updates/adds a StringKeyMapTypeSection in StringKeyMapTypeSections.
func (s StringKeyMapTypeSections) UpdateElemToSections(updated StringKeyMapTypeSection) {
	s.RemoveElemInSections(updated.ParentKey, updated.NoOfArrayLv)
	s = append(s, updated)
}

// RemoveElemInSections removes a StringKeyMapTypeSection from StringKeyMapTypeSections and assigns the new slice to StringKeyMapTypeSections.
func (s StringKeyMapTypeSections) RemoveElemInSections(key string, noOfArrayLv int) {
	y := []StringKeyMapTypeSection{}
	for _, x := range s {
		if !(x.ParentKey == key && x.NoOfArrayLv == noOfArrayLv) {
			y = append(y, x)
		}
	}
	s = y
}

// CheckIfValueTypeAlreadyKnown finds out if the value's type is known. index is used to locate the value in an array.
func (s StringKeyMapTypeSections) CheckIfValueTypeAlreadyKnown(key string, noOfArrayLv int) (valueType reflect.Kind, known bool) {
	var isInArray bool
	if noOfArrayLv == 0 {
		isInArray = false
	} else {
		isInArray = true
	}
	for _, x := range s {
		if key == x.ParentKey {
			if isInArray {
				if x.NoOfArrayLv == noOfArrayLv {
					return x.TypeSlice, true
				}
			} else {
				return x.TypeMap[key], true
			}
		}
	}
	return reflect.Invalid, false
}

// ExploreUnknownSection tries to find out the section detail for a unknown map/slice value.
func ExploreUnknownSection(key string, v reflect.Value, noOfArrayLv int) (sec StringKeyMapTypeSection, discovered bool) {
	switch v.Kind() {
	case reflect.Map:
		sc, isMap := FindValueTypesInUnknownMap(key, v, noOfArrayLv)
		if isMap {
			return sc, true
		}
	case reflect.Slice:
		sc, isSlice := FindValueTypesInSlice(key, v, noOfArrayLv)
		if isSlice {
			return sc, true
		}
	default:
		fmt.Println("Not a map/slice.")
	}
	return StringKeyMapTypeSection{}, false
}
