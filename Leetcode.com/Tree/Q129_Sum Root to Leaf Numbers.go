package Tree

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	sumTree(root,&sum,nil)
	return sum
}

//切片是将地址复制过去，但是在代码里面append会导致数组变换地址，导致不符合原来的预期
func sumTree(root *TreeNode, sum *int, nums []int) {

	if root.Left == nil && root.Right == nil {
		tmp := 0
		for _, num := range nums {
			tmp = tmp * 10 + num
		}
		tmp = tmp * 10 + root.Val
		*sum += tmp
	}

	if root.Left != nil {
		sumTree(root.Left,sum, append(nums,root.Val))
	}

	if root.Right != nil {
		sumTree(root.Right, sum, append(nums,root.Val))
	}
}