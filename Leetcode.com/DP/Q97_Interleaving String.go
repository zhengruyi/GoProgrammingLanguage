package DP

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 + l2 != len(s3) {
		return false
	}
	// go 没办法初始化多维数组，只能这样挨个初始化
	dp := make([][]bool,l1+1)
	for i := range dp {
		dp[i] = make([]bool,l2+1)
	}

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			//都为空，视为匹配上
			if i == 0 && j == 0 {
				dp[i][j] = true
			//分别表示单边匹配
			}else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j -1]
			}else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			//i，j都不为0时，只要i，j中的一个匹配上就可以了
			}else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dp[l1][l2]
}