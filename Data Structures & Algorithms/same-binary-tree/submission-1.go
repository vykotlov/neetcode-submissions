func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	for root != nil {
		if root.Val > subRoot.Val {
			root = root.Right
		}

		if root.Val < subRoot.Val {
			root = root.Left
		}

		return isSubtree(root, subRoot)
	}

	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
