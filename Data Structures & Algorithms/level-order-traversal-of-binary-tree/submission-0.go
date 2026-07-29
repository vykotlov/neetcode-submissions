func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	allLevels := [][]int{}
	currentLevel := []int{}
	
	currentNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	
	for len(currentNodes) > 0 {
		for _, node := range currentNodes {
			currentLevel = append(currentLevel, node.Val)
			
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}
			
			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		
		allLevels = append(allLevels, currentLevel)
		currentLevel = []int{}
		
		currentNodes = nextNodes
		nextNodes = []*TreeNode{}
	}
	
	return allLevels
}
