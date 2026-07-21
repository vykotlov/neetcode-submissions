func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node == nil {
			return
		}
		
		recursive(node.Left)
		recursive(node.Right)
		result = append(result, node.Val)
	}

	recursive(root)
	return result
}