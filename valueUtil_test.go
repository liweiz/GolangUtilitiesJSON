package utilities

import "testing"

type testCompareValuesInOut struct {
	Ins []interface{}
	Out bool
}

func TestCompareValues(t *testing.T) {
	toTest := []testCompareValuesInOut{
		testCompareValuesInOut{
			[]interface{}{nil, m9},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{m9, m9},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{m9, m1},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{true, true},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{true, false},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{0.1, 0.1},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{0.1, 0.2},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{"m9", "m9"},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{"m9", ""},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{s1, s1},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{s1, s2},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1, m3}},
			true,
		},
		testCompareValuesInOut{
			[]interface{}{[]map[string]interface{}{m1, m3}, []map[string]interface{}{m1}},
			false,
		},
		testCompareValuesInOut{
			[]interface{}{[]float64{0.8, 0.9}, []float64{0.8, 0.9}},
			false,
		},
	}
	for _, x := range toTest {
		if CompareValues(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: interface{}: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}
