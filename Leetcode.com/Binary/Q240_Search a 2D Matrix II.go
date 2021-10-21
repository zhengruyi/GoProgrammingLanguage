package Binary

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	i := 0
	j := cols - 1
	for ; i < rows && j >= 0; {
		if matrix[i][j] == target {
			return true
		}else if matrix[i][j] > target {
			j--
		}else if matrix[i][j] < target{
			i++
		}
	}
	return false
}
