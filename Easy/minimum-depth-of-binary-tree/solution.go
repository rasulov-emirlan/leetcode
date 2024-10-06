package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := 0
	rightCount := 0

	if root.Left != nil {
		leftCount += minDepth(root.Left) + 1
	}

	if root.Right != nil {
		rightCount += minDepth(root.Right) + 1
	}

	if leftCount != 0 && rightCount != 0 {
		if leftCount < rightCount {
			return leftCount
		}
		return rightCount
	}

	if rightCount != 0 {
		return rightCount
	}
	if leftCount != 0 {
		return leftCount
	}

	return 1
}
