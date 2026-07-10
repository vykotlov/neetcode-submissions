func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodes := []*TreeNode{root}
	newNodes := []*TreeNode{}
	depth := 0

	for len(nodes) > 0 {
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}

		depth++
		nodes = newNodes
		newNodes = []*TreeNode{}
	}

	return depth
}
