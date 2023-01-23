package balancedbinarytree

import "testing"

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		desc string
		root *TreeNode
		want bool
	}{
		{
			desc: "Example 1",
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
			want: true,
		},
		{
			desc: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := isBalanced(tc.root)
			if got != tc.want {
				t.Errorf("isBalanced(%v) = %v; want %v", tc.root, got, tc.want)
			}
		})
	}
}
