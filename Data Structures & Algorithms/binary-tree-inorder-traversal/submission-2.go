func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		recursive(node.Left)
		result = append(result, node.Val)
		recursive(node.Right)
	}

	recursive(root)
	return result
}