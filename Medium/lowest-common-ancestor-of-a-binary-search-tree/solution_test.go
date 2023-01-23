package lowestcommonancestorofabinarysearchtree

import "testing"

func TestLowestAncestor(t *testing.T) {
	testCases := []struct {
		desc string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			desc: "Example 1",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			p: &TreeNode{
				Val: 2,
			},
			q: &TreeNode{
				Val: 8,
			},
			want: &TreeNode{
				Val: 6,
			},
		},
		{
			desc: "Example 2",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			p: &TreeNode{
				Val: 2,
			},
			q: &TreeNode{
				Val: 4,
			},
			want: &TreeNode{
				Val: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			if got == nil {
				t.Fatalf("got nil, want %v", tc.want)
			}

			if got.Val != tc.want.Val {
				t.Fatalf("got %d, want %d", got.Val, tc.want.Val)
			}
		})
	}
}
