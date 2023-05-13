package maximumdepthofbinarytree

import "testing"

func TestMaximumDepth(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
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
			want: 3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: 2,
		},
		{
			root: &TreeNode{
				Val: 0,
			},
			want: 1,
		},
	}

	for _, c := range cases {
		got := maxDepth(c.root)
		if got != c.want {
			t.Errorf("maxDepth(%v) == %d, want %d", c.root, got, c.want)
		}
	}
}
