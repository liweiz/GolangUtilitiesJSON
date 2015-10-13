package utilities

import (
	"fmt"
	"reflect"
)

// StringKeyMapTypeMap stores StringKeyMap's value type for one layer of one branch.
type StringKeyMapTypeMap map[string]reflect.Kind

// StringKeyMapTypeSection stores the parent key and possible array index to link the StringKeyMapType info for its parent layer. An empty ParentKey indicates it's a root map. NoOfArrayLv == -1 indicates parent level is not an array.
type StringKeyMapTypeSection struct {
	ParentKey   string
	NoOfArrayLv int
	TypeMap     map[string]reflect.Kind
	TypeSlice   reflect.Kind
}

// ValueForKeyInTypeMap finds the type of value for a key.
func (s StringKeyMapTypeSection) ValueForKeyInTypeMap(key string) (kind reflect.Kind, found bool) {
	if t, ok := s.TypeMap[key]; ok {
		if t != reflect.Invalid {
			return t, true
		}
	}
	return reflect.Invalid, false
}

// FindValueTypesInKnownMap finds out all value types in this level in a known map.
func (s StringKeyMapTypeSections) FindValueTypesInKnownMap(key string, v reflect.Value, noOfArrayLv int) (section StringKeyMapTypeSection, isMap bool) {
	if v.Kind() == reflect.Map {
		i := reflect.Invalid
		if len(v.MapKeys()) > 0 {
			i = reflect.Map
		}
		r := StringKeyMapTypeSection{key, noOfArrayLv, StringKeyMapTypeMap{}, i}
		newTypeFound := false
		for _, x := range v.MapKeys() {
			t, known := s.CheckIfValueTypeAlreadyKnown(x.String(), noOfArrayLv)
			if known {
				r.TypeMap[x.String()] = t
			} else {
				k := FindTypeForValue(v.MapIndex(x))
				if k == reflect.Invalid {
					return StringKeyMapTypeSection{}, false
				}
				r.TypeMap[x.String()] = k
				newTypeFound = true
			}
		}
		if newTypeFound {
			s.RemoveElemInSections(key, noOfArrayLv)
		}
		return r, true
	}
	return StringKeyMapTypeSection{}, false
}

// FindValueTypesInUnknownMap finds out all value types in this level in a unknown map.
func FindValueTypesInUnknownMap(key string, v reflect.Value, noOfArrayLv int) (section StringKeyMapTypeSection, isMap bool) {
	if v.Kind() == reflect.Map {
		if v.Len() > 0 {
			s := StringKeyMapTypeSection{}
			s.TypeMap = StringKeyMapTypeMap{}
			for _, x := range v.MapKeys() {
				fmt.Println("FindValueTypesInUnknownMap xxxxxx", v.MapIndex(x))
				k := FindTypeForValue(v.MapIndex(x))
				fmt.Println("FindValueTypesInUnknownMap yyyyyy: ", k.String())
				s.TypeMap[x.String()] = k
				fmt.Println("FindValueTypesInUnknownMap zzzzzz: ", s.TypeMap[x.String()])
			}
			s.ParentKey = key
			s.NoOfArrayLv = noOfArrayLv
			return s, true
		}
	}
	return StringKeyMapTypeSection{}, false
}

// FindValueTypesInSlice finds out all value types in this level in a slice.
func FindValueTypesInSlice(key string, v reflect.Value, noOfArrayLv int) (section StringKeyMapTypeSection, isSlice bool) {
	if v.Kind() == reflect.Slice {
		if v.Len() > 0 {
			s, _ := ConvertToValueSlice(v)
			k := FindTypeForValue(s[0])
			return StringKeyMapTypeSection{key, noOfArrayLv, StringKeyMapTypeMap{}, k}, true
		}
	}
	return StringKeyMapTypeSection{}, false
}
