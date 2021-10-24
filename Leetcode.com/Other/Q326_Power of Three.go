package Other

func isPowerOfThree(n int) bool {
	// 1162261467是底数位3,不溢出的最大次方数字
	return (n > 0 && 1162261467 % n == 0)
}
