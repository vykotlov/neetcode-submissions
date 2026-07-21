func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		result = append(result, node.Val)
		recursive(node.Left)
		recursive(node.Right)
	}

	recursive(root)
	return result
}