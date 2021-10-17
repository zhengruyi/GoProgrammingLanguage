package BFS

//遍历完后两两进行交换
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue [] *TreeNode
	queue = append(queue,root)
	for; len(queue) > 0; {
		length := len(queue)
		var tmp []int
		for j := 0; j < length; j++ {
			tmp = append(tmp,queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue,queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue,queue[0].Right)
			}
			queue = queue[1:]
		}
		res = append(res,tmp)
	}
	for l, r := 0, len(res) -1; l < r;{
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}