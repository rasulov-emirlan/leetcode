package sumofleftleaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return recursive(root.Left, true) + recursive(root.Right, false)
}

func recursive(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return recursive(root.Left, true) + recursive(root.Right, false)
}
