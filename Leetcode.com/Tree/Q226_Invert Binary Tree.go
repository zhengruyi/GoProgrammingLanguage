package Tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left,root.Right = root.Right, root.Left

	if root.Left != nil {
		invert(root.Left)
	}
	if root.Right != nil {
		invert(root.Right)
	}
}
