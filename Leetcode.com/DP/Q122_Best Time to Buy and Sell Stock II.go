package DP

func maxProfit(prices []int) int {
	sum := 0
	// 可以当天买卖，所以就是多段利润之和
	for i := 0; i < len(prices) - 1; i++ {
		if prices[i+1] - prices[i] > 0 {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
