package symmetrictree

// solved with a lot of help from solutions on leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return "🌳"
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
