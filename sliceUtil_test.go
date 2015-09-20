package utilities

import "testing"

var s1 = []string{"a", "b", "c"}
var s2 = []string{"a"}
var s3 = []string{"b", "c", "a"}
var s4 = []string{"a", "d", "c"}
var s5 = []string{"a", "b", "c"}
var m1 = map[string]interface{}{"string": "a"}
var m2 = map[string]interface{}{"bool": true}
var m3 = map[string]interface{}{"float64": float64(45)}
var m4 = map[string]interface{}{"[]string": []string{"a", "b", "c"}}
var m5 = map[string]interface{}{"map": m1}
var m6 = map[string]interface{}{"string": "b"}
var m7 = map[string]interface{}{"bool": false}
var m8 = map[string]interface{}{"float64": float64(43)}
var m9 = map[string]interface{}{"[]string": []string{"d", "b", "c"}}
var m10 = map[string]interface{}{"map": m6}
var m11 = map[string]interface{}{"bool": true}
var m12 = map[string]interface{}{"bool": true, "map": m6}

type testStringSliceInBoolOut struct {
	Ins [][]string
	Out bool
}

type testStringKeyMapSliceInBoolOut struct {
	InInterface    []interface{}
	InNonInterface []map[string]interface{}
	Out            bool
}

type testStringKeyMapSlicesInBoolOut struct {
	Ins [][]map[string]interface{}
	Out bool
}

func TestCompareStringKeyMapSlices(t *testing.T) {
	toTest := []testStringKeyMapSlicesInBoolOut{
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m1, m10}},
			true,
		},
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{m1, m10}, []map[string]interface{}{m10}},
			false,
		},
		testStringKeyMapSlicesInBoolOut{
			[][]map[string]interface{}{[]map[string]interface{}{}, []map[string]interface{}{m10}},
			false,
		},
	}
	for _, x := range toTest {
		if CompareStringKeyMapSlices(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: Slices: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}

func TestCheckIfIsStringKeyMapSlice(t *testing.T) {
	toTest := []testStringKeyMapSliceInBoolOut{
		testStringKeyMapSliceInBoolOut{
			[]interface{}{m1, m10},
			[]map[string]interface{}{m1, m10},
			true,
		},
		testStringKeyMapSliceInBoolOut{
			[]interface{}{},
			[]map[string]interface{}{},
			false,
		},
	}
	for _, x := range toTest {
		r, s := CheckIfIsStringKeyMapSlice(x.InInterface)
		if r != x.Out {
			t.Errorf("ERR: Slices: %+v CheckIfIsStringKeyMapSlice test failed. Should be %+v\n", x.InInterface, x.Out)
		} else if x.Out {
			var ok = true
			if len(s) == len(x.InNonInterface) {
				for i, y := range s {
					if !CompareValues(y, x.InNonInterface[i]) {
						ok = false
						break
					}
				}
			} else {
				ok = false
			}
			if !ok {
				t.Errorf("ERR: Slices: %+v CheckIfIsStringKeyMapSlice test failed. Output slice should be %+v\n", x.InInterface, x.InNonInterface)
			}
		}
	}
}

func TestCompareStringSlices(t *testing.T) {
	toTest := []testStringSliceInBoolOut{
		testStringSliceInBoolOut{
			[][]string{s1, s2},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s3},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s4},
			false,
		},
		testStringSliceInBoolOut{
			[][]string{s1, s5},
			true,
		},
	}
	for _, x := range toTest {
		if CompareStringSlices(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: Slices: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}

// func TestCompareMapValues(t *testing.T) {
// 	assert.True(t, CompareMapValues(m1["string"], m1["string"]), "ERR: string value in map true test failed.")
// 	assert.True(t, CompareMapValues(m2["bool"], m2["bool"]), "ERR: bool value in map true test failed.")
// 	assert.True(t, CompareMapValues(m3["float64"], m3["float64"]), "ERR: float64 value in map true test failed.")
// 	assert.True(t, CompareMapValues(m4["[]string"], m4["[]string"]), "ERR: []string value in map true test failed.")
// 	assert.True(t, CompareMapValues(m5["map"], m5["map"]), "ERR: map value in map true test failed.")
// 	assert.False(t, CompareMapValues(m1["string"], m6["string"]), "ERR: string value in map false test failed.")
// 	assert.False(t, CompareMapValues(m2["bool"], m7["bool"]), "ERR: bool value in map false test failed.")
// 	assert.False(t, CompareMapValues(m3["float64"], m8["float64"]), "ERR: float64 value in map false test failed.")
// 	assert.False(t, CompareMapValues(m4["[]string"], m9["[]string"]), "ERR: []string value in map false test failed.")
// 	assert.False(t, CompareMapValues(m5["map"], m10["map"]), "ERR: map value in map false test failed.")
// }
