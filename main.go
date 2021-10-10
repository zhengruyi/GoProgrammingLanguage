package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	sumTree(root, &sum, nil)
	return sum
}

func sumTree(root *TreeNode, sum *int, nums []int) {

	if root.Left == nil && root.Right == nil {
		tmp := 0
		for _, num := range nums {
			tmp = tmp*10 + num
		}
		*sum += tmp
	}

	if root.Left != nil {
		sumTree(root.Left, sum, append(nums, root.Val))
	}

	if root.Right != nil {
		sumTree(root.Right, sum, append(nums, root.Val))
	}
}
func main() {
	left := &TreeNode{2,nil,nil}
	right := &TreeNode{3,nil,nil}
	root := &TreeNode{1,left,right}
	sumNumbers(root)
}
