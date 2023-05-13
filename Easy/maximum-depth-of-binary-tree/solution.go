package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return "🌳"
}

// terrible solution but whatever

func maxDepth(root *TreeNode) int {
	sum := check(root)
	if sum == 0 && root == nil {
		return 0
	}
	return sum + 1
}

func check(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	if root.Left != nil || root.Right != nil {
		sum++

		suml := 0
		sumr := 0

		if root.Left != nil {
			suml = check(root.Left)
		}
		if root.Right != nil {
			sumr = check(root.Right)
		}

		if suml > sumr {
			sum += suml
		} else {
			sum += sumr
		}
	}

	return sum
}
