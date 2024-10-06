package minimumdepthofbinarytree

import "testing"

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
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
			expected: 2,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := minDepth(test.root)
			if result != test.expected {
				t.Errorf("Expected minimum depth to be %d, but got %d", test.expected, result)
			}
		})
	}
}
