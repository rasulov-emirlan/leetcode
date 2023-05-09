package sametree

import "testing"

func TestSameTree(t *testing.T) {
	cases := []struct {
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},

			expected: false,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			expected: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		res := isSameTree(c.p, c.q)
		if res != c.expected {
			t.Errorf(
				"\ninput p: %v \ninput q: %v \nexpected: %v \nactual: %v",
				c.p,
				c.q,
				c.expected,
				res,
			)
		}
	}
}
