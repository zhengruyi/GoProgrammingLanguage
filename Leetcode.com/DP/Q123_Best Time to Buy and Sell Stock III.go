package DP

// 具体分析：https://labuladong.gitbook.io/algo/mu-lu-ye-2/mu-lu-ye-4/tuan-mie-gu-piao-wen-ti
func maxProfitStock(prices []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	dp_20 := 0
	dp_21 := INT_MIN
	dp_10 := 0
	dp_11 := INT_MIN
	for i := range prices {
		dp_20 = max(dp_20, dp_21+prices[i])
		dp_21 = max(dp_21, dp_10-prices[i])
		dp_10 = max(dp_10, dp_11+prices[i])
		dp_11 = max(dp_11, -prices[i])

	}
	return dp_20

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
