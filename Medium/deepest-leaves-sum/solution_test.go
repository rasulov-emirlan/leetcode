package deepestleavessum

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			want: 15,
		},
	}

	for _, c := range cases {
		got := deepestLeavesSum(c.root)

		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
