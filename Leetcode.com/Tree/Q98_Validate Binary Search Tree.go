package Tree

func isValidBST(root *TreeNode) bool {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	return isBST(root,INT_MIN,INT_MAX)
}
// 决定左右边界，时时更新
func isBST(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isBST(root.Left,min,root.Val) && isBST(root.Right,root.Val,max)
}
