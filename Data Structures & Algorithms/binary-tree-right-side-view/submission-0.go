func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	current := []*TreeNode{root}
	next := []*TreeNode{}
	result := []int{}

	for len(current) > 0 {
		for _, node := range current {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		
		result = append(result, current[len(current)-1].Val)
		current = next
		next = []*TreeNode{}
	}
	
	return result
}
