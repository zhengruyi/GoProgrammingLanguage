package Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root != nil {
		// ...表示任意多个参数，所以可以将【】int解释成多个元素，挨个添加进切片中
		ans = append(ans, inorderTraversal(root.Left)...)
		ans = append(ans, root.Val)
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}
