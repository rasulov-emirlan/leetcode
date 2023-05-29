package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			output: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for i, c := range cases {
		rotate(c.input)
		if !reflect.DeepEqual(c.input, c.output) {
			t.Errorf("case %d: want %v got %v", i, c.output, c.input)
		}
	}
}
