package DP

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			//0..j之间的字符串可以有word dict中的字符串组成，那么如果j..i也可以在dict中找到，那么0..i全部可以在砸店中找到
			if dp[j] && contains(wordDict, string(s[j:i])) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 检测列表中是否包含目标元素
func contains(strs []string, target string) bool{
	for _, v := range strs {
		if v == target {
			return true
		}
	}
	return false
}