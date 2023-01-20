package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	return recursive(root)
}

// Accepted. Beats 100% in speed. Beats 100% in memory. Memory usage is 2.1 MB.
func recursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = recursive(root.Right), recursive(root.Left)

	return root
}
