package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		size := len(queue)
		leaves := []int{}
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			leaves = append(leaves, curr.Val)
		}
		result = append(result, leaves)
	}

	return result
}
