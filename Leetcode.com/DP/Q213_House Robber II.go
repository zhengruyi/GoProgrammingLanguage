package DP


//因为首尾相连的原因,所以这个问题转化成两个问题,0..n-1和1..n.两种方式的最大值就是答案
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return maxII(dp(nums[0:len(nums)-1]),dp(nums[1:len(nums)]))
}

func dp (nums []int) int {
	dp := make([]int,len(nums) + 2)
	for i := 2; i < len(nums) + 2; i++ {
		dp[i] = maxII(dp[i-2]+nums[i-2],dp[i-1])
	}
	return dp[len(nums) + 1]
}

func maxII(a, b int) int{
	if a > b {
		return a
	}
	return b
}