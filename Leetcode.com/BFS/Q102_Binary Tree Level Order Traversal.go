package BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
// BFS 遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for ; len(queue) > 0; {
		length := len(queue)
		var tmp []int
		for j := 0; j < length; j++ {
			tmp = append(tmp, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res, tmp)
	}
	return res
}
