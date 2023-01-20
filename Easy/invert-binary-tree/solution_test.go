package invertbinarytree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			input: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			expected: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "Example 2",
			input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expected: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := invertTree(tc.input)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}

}
