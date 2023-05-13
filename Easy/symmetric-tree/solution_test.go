package symmetrictree

import "testing"

func TestIsSimmetric(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			want: false,
		},
	}

	for _, c := range cases {
		got := isSymmetric(c.root)
		if got != c.want {
			t.Errorf("isSymmetric(%v) == %t, want %t", c.root, got, c.want)
		}
	}
}
