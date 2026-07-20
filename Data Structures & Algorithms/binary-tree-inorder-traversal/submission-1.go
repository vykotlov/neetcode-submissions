func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	
	recursive := func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}

	recursive(root)
	return result
}
