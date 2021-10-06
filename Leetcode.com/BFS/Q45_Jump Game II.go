package BFS

// BFS思想，每次可以便利的范围是[i,lastJump]范围，遍历完就需要向下一层便利，这时count++
func jump(nums []int) int {
	currentJump := 0
	lastJump := 0
	count := 0
	for i := 0; i < len(nums) - 1; i++{
		currentJump = max(currentJump,i + nums[i])
		if i == lastJump {
			count ++
			lastJump = currentJump
		}
	}
	return count
}

func max(a,b int)int {
	if a > b {
		return a
	}
	return b
}