package combinationsum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var tests = []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
	}

	for _, test := range tests {
		if output := combinationSum(test.candidates, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("combinationSum(%v, %d) = %v but expected %v", test.candidates, test.target, output, test.expected)
		}
	}
}
