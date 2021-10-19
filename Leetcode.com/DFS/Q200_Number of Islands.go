package DFS

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs200(grid,i,j)
				res++
			}
		}
	}
	return res
}

func dfs200(grid [][]byte, i,j int) {
	grid[i][j] = '0'
	direct := [][]int{{0,-1},{0,1},{-1,0},{1,0}}
	rows := len(grid)
	cols := len(grid[0])
	for _, offset := range direct {
		newX := i + offset[0]
		newY := j + offset[1]
		if newX >= 0 && newX < rows && newY >=0 && newY < cols && grid[newX][newY] == '1' {
			dfs(grid,newX,newY)
		}
	}
}
