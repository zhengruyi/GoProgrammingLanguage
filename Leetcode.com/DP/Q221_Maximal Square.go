package DP

func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])
	dp := make([][]int,row+1)
	for i, _ := range dp {
		dp[i] = make([]int, col+1)
	}
	res := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if matrix[i-1][j-1] == '1' {
				//正方形的边长取决于三个点的最小值
				dp[i][j] = min(dp[i-1][j-1],min(dp[i-1][j],dp[i][j-1])) + 1
				res = maxSquare(dp[i][j],res)
			}
		}
	}
	return res * res
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
func maxSquare(a,b int) int {
	if a > b {
		return a
	}
	return b
}
