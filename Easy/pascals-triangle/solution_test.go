package pascalstriangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		numRows int
		want    [][]int
	}{
		{0, nil},
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	}

	for _, c := range cases {
		got := generate(c.numRows)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("generate(%d) == %v, want %v", c.numRows, got, c.want)
		}
	}
}
