package DP

func maxProfitStockI(prices []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	dp_10 := 0
	dp_11 := INT_MIN
	for i :=  range prices {
		dp_10 = max(dp_10, dp_11 + prices[i])
		dp_11 = max(dp_11, -prices[i])
	}
	return dp_10
}

func maxI(a,b int) int {
	if a > b {
		return a
	}
	return b
}