package BFS

func canJump(nums []int) bool {
	currentJump :=0
	for i := 0; i < len(nums); i++ {
		// 如果前面所有节点的最大距离都不能到达当前节点，那么就反悔false
		if i > currentJump {
			return false
		}
		currentJump = max55(currentJump,i+nums[i])
	}
	return true
}

func max55(a,b int) int {
	if a > b {
		return a
	}
	return b
}