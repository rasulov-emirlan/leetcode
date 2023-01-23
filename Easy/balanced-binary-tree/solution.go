package balancedbinarytree

// Accepted

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := height(root.Left)
	right := height(root.Right)

	if abs(left-right) <= 1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right) {
		return true
	}

	return false
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
