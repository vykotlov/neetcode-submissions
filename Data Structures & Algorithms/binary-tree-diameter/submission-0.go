func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := diameterOfBinaryTreeRecursive(root)

	return diameter
}

func diameterOfBinaryTreeRecursive(root *TreeNode) (depth int, diameter int) {
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := diameterOfBinaryTreeRecursive(root.Left)
	rightDepth, rightDiameter := diameterOfBinaryTreeRecursive(root.Right)

	currentDiameter := leftDepth + rightDepth
	maxDiameter := max(currentDiameter, leftDiameter, rightDiameter)
	maxDepth := max(leftDepth, rightDepth)

	return maxDepth + 1, maxDiameter
}
