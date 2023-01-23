package lowestcommonancestorofabinarysearchtree

// Accepted

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	lca := root

	for lca != nil {
		if lca.Val < p.Val && lca.Val < q.Val {
			lca = lca.Right
			continue
		}

		if lca.Val > p.Val && lca.Val > q.Val {
			lca = lca.Left
			continue
		}

		return lca
	}

	return lca
}
