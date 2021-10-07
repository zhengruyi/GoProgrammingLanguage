package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var res []int
	sumTree(root, 0, res)
	sum := 0
	for _, num := range res {
		sum += num
	}
	return sum
}

func sumTree(root *TreeNode, num int, res []int) []int{
	if root != nil {
		num = num * 10 + root.Val
	}

	if root.Left == nil && root.Right == nil {
		res = append(res,num)
		return res
	}

	if root.Left != nil {
		res = append(res,sumTree(root.Left, num,res)...)
	}
	if root.Right != nil {
		res = append(res,sumTree(root.Right,num, res)...)
	}
	return res
}

func main() {
	left := &TreeNode{2,nil,nil}
	right := &TreeNode{3,nil,nil}
	root := &TreeNode{1,left,right}
	sumNumbers(root)
}
