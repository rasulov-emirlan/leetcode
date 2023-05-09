package binarytreeinordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
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
			want: []int{1, 3, 2},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
		{
			root: nil,
			want: []int{},
		},
	}

	for _, c := range cases {
		got := inorderTraversal(c.root)
		if len(got) != len(c.want) {
			t.Errorf("inorderTraversal(%v) == %v, want %v", c.root, got, c.want)
		}
		if len(c.want)+len(got) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("inorderTraversal(%v) == %v, want %v", c.root, got, c.want)
		}
	}
}
