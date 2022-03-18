package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		answer []int
	}{
		{
			arr:    []int{2, 7, 11, 15},
			target: 9,
			answer: []int{0, 1},
		},
		{
			arr:    []int{3, 2, 4},
			target: 6,
			answer: []int{1, 2},
		},
		{
			arr:    []int{3, 3},
			target: 6,
			answer: []int{0, 1},
		},
	}

	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			answer := twoSum(tC.arr, tC.target)
			if answer[0] != tC.answer[0] || answer[1] != tC.answer[1] {
				t.Error("incorrect")
			}
		})
	}
}
