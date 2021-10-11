package DP

func rob(nums []int) int {
	dp := make([]int, len(nums) + 2)

	for i := 2; i < len(nums) + 2; i++ {
		//到当前这家的最大利润等于抢当前这家+往前两家的最大利润,或者不抢,等于到上一家的利润
		dp[i] = maxRob(dp[i-1],dp[i-2] + nums[i-2])
	}
	return dp[len(nums) + 1]
}

func maxRob (a,b int) int {
	if a > b {
		return a
	}
	return b
}
