package binarytreepreordertraversal

import (
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 2, 3},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
	}

	for _, c := range cases {
		got := preorderTraversal(c.root)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("preorderTraversal(%v) == %v, want %v", c.root, got, c.want)
		}
	}
}
