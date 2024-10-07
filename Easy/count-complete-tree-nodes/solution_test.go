package countcompletetreenodes

import "testing"

func TestCountNodes(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,

					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		if res := countNodes(tc.root); res != tc.expected {
			t.Fatalf("expected %d, got %d", tc.expected, res)
		}
	}
}
