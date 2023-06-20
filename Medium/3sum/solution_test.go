package sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums   []int
		result [][]int
	}{
		{
			nums:   []int{-1, 0, 1, 2, -1, -4},
			result: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:   []int{0, 0, 0, 0},
			result: [][]int{{0, 0, 0}},
		},
		{
			nums:   []int{-2, 0, 1, 1, 2},
			result: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}

	for _, c := range cases {
		if res := threeSum(c.nums); !compare(res, c.result) {
			t.Errorf("threeSum(%v)=%v, expected %v", c.nums, res, c.result)
		}
	}
}

func compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		sort.Ints(a[i])
		sort.Ints(b[i])
	}

	return reflect.DeepEqual(a, b)
}
