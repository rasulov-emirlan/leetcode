package binarytreelevelordertraversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	cases := []struct {
		root   *TreeNode
		expect [][]int
	}{
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}

	for _, c := range cases {
		if result := levelOrder(c.root); !reflect.DeepEqual(result, c.expect) {
			t.Fatalf("levelOrder(%v) fails.\nExpected %v\nReceived %v", c.root, c.expect, result)
		}
	}
}
