package Other

func countPrimes(n int) int {
	count := 0
	stats := make([]bool,n)
	// 一个数字无论他是不是质数,他的所有倍数都不在是质数,所以通过这种方法可以找到所有质数
	for i := 2; i < n; i++ {
		if !stats[i] {
			count ++
			for j := 2; j * i < n; j++ {
				stats[j*i] = true
			}
		}
	}
	return count
}