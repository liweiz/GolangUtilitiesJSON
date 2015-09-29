package utilities

import "reflect"

// ValuePath is the container for a level-based path to a value in a StringKeyMap. It also contains single and slice (if applicable) forms of value(s).
// Level here is not just identified by keys in different levels. It's also identified by slices.
// Keys is a slice to record the keys for all levels on the path. If the value is in an array, the corresponding key for that level is empty string: "" since empty string key in JSON seems have no real value. We can simply ignore it here.
// Indexes is a slice to record the index for all levels on the path. If the value is in a map, the corresponding index for that level is -1.
// ReflectValue is the corresponding value stored as a single value.
type ValuePath struct {
	Keys         []string
	Indexes      []int16
	ReflectValue reflect.Value
}

// ValueKey returns the key for the value.
func (p ValuePath) ValueKey() string {
	l := len(p.Keys)
	if l > 1 {
		if len(p.Keys[l-1]) == 0 {
			return p.Keys[l-2]
		}
		return p.Keys[l-1]
	}
	return p.Keys[0]
}

// ValueIndex returns the index for the value, if the value is in an array. Otherwise, it returns -1 to indicate the value is not from an array.
func (p ValuePath) ValueIndex() int16 {
	l := len(p.Keys)
	if l > 1 {
		if len(p.Keys[l-1]) == 0 {
			return p.Indexes[l-1]
		}
	}
	return -1
}
