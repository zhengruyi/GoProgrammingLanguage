package Array

func lengthOfLongestSubstring(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	// 统计出现的字符和出现的位置
	stas := make(map[string]int)
	left := 0
	index := 0
	var char rune
	for index, char = range s {
		//如果出现了重复字符，那么就缩减左边界
		appear, ok := stas[string(char)]
		if ok {
			//将所有出现的记录全部删掉
			for i := left; left <= appear; left++ {
				delete(stas, string(s[i]))
			}
		}
		//每次都计算一次最长无重复字串
		res = max(res, index-left+1)
		//更新字符和出现位置
		stas[string(char)] = index
	}
	//如果全程无重复字符串，那么这里触发最终的一次计算
	res = max(res, index-left+1)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
