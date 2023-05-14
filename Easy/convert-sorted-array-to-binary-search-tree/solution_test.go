package convertsortedarraytobinarysearchtree

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	cases := []struct {
		nums   []int
		expect *TreeNode
	}{
		{
			nums: []int{-10, -3, 0, 5, 9},
			expect: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -10,
					Right: &TreeNode{
						Val: -3,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
		},
	}

	for _, c := range cases {
		if actual := sortedArrayToBST(c.nums); !reflect.DeepEqual(actual, c.expect) {
			t.Errorf("sortedArrayToBST(%v)=%v, expected %v", c.nums, actual, c.expect)
		}
	}
}
