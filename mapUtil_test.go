package utilities

import (
	"fmt"
	"testing"
)

type testCompareStringKeyMapsInOut struct {
	Ins []map[string]interface{}
	Out bool
}

func TestCompareStringKeyMaps(t *testing.T) {
	fmt.Println("\nSTART TestCompareStringKeyMaps")
	toTest := []testCompareStringKeyMapsInOut{
		testCompareStringKeyMapsInOut{
			[]map[string]interface{}{m9, m9},
			true,
		},
		testCompareStringKeyMapsInOut{
			[]map[string]interface{}{m7, m11},
			false,
		},
		testCompareStringKeyMapsInOut{
			[]map[string]interface{}{m4, m9},
			false,
		},
		testCompareStringKeyMapsInOut{
			[]map[string]interface{}{m10, m12},
			false,
		},
	}
	for _, x := range toTest {
		if CompareStringKeyMaps(x.Ins[0], x.Ins[1]) != x.Out {
			t.Errorf("ERR: StringKeyMaps: %+v AND %+v comparison failed. Should be %+v\n", x.Ins[0], x.Ins[1], x.Out)
		}
	}
}
