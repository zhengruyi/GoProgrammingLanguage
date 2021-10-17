package BFS

// BFS,每次遍历完一层,高度就加1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue [] *TreeNode
	depth := 0
	queue = append(queue,root)
	for ; len(queue) > 0; {
		length := len(queue)
		for j := 0; j < length; j++ {
			if queue[0].Left != nil {
				queue = append(queue,queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue,queue[0].Right)
			}
			queue = queue[1:]
		}
		depth++
	}
	return depth
}
