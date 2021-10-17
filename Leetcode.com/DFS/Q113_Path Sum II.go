package DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//返回从根节点返回的序列
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			res = append(res, []int{root.Val})
		}
		return res
	}
	//从返回的序列中,加入当前节点进行返回
	for _, path := range pathSum(root.Left, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	for _, path := range pathSum(root.Right, targetSum-root.Val) {
		res = append(res, append([]int{root.Val}, path...))
	}
	return res
}
