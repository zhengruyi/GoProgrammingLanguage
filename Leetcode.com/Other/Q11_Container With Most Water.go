package Other

// 盛水的容器这题要明白随着两者之间长度的减小，只有高度增大才有可能导致容量增大
// 而容器的高度取决于两个高度的最小值，所以优先移动较低的容器
func maxArea(height []int) int {
	res := 0
	i := 0
	j := len(height) - 1
	for ; i < j; {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
