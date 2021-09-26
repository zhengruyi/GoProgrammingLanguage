package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	recursive(root, res)
	return res
}
func recursive(root *TreeNode, res []int) {
	if root == nil {
		return
	}
	recursive(root.Left, res)
	res = append(res, root.Val)
	recursive(root.Right, res)
}


func main() {
	left := &TreeNode{1,nil,nil}
	right := &TreeNode{3,nil,nil}
	root := &TreeNode{2,left,right}
	inorderTraversal(root)
}
