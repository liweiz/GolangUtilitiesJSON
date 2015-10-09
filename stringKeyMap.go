package utilities

// StringKeyMap is the map counterpart for JSON.
type StringKeyMap map[string]interface{}

// Keys return a slice of all keys.
func (s StringKeyMap) Keys() []string {
	keys := make([]string, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i++
	}
	return keys
}
