package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{2, 5}, 5, 1},
		{[]int{2, 5}, 2, 0},
		{[]int{2, 5}, 1, -1},
		{[]int{2, 5}, 6, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v find %d", tc.nums, tc.target), func(t *testing.T) {
			got := search(tc.nums, tc.target)
			if got != tc.expect {
				t.Errorf("got %d, expect %d", got, tc.expect)
			}
		})
	}
}
