package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			result: []int{1},
		},
	}

	for _, c := range cases {
		merge(c.nums1, c.m, c.nums2, c.n)
		if !reflect.DeepEqual(c.nums1, c.result) {
			t.Errorf(
				"merge(%v, %d, %v, %d) == %v, want %v",
				c.nums1,
				c.m,
				c.nums2,
				c.n,
				c.nums1,
				c.result,
			)
		}
	}
}
