package Other

// 用字典来做set
func longestConsecutive(nums []int) int {
	set := make(map[int]int)
	for _, num := range nums {
		set[num] = 0
	}
	res := 0
	for _, num := range nums {
		// 如果当前数字是联塑数字的起点,那么选当前数字为起点,然后逐步往前走
		if _,ok := set[num - 1]; !ok {
			next := num + 1
			for;; {
				//计算以num为起点,然后往前遍历,找到最长序列
				if _, ok := set[next]; ok {
					next += 1
				}else{
					break
				}
			}
			res = max128(res,next - num)
		}
	}
	return res
}

func max128(a,b int) int {
	if a > b {
		return a
	}
	return b
}
