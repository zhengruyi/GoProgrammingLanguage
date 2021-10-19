package main

import "fmt"

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byte('1') {
				dfs(grid,i,j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i,j int) {
	grid[i][j] = byte('0')
	direct := [][]int{{1,-1},{1,1},{-1,1},{-1,-1}}
	rows := len(grid)
	cols := len(grid[0])
	for _, offset := range direct {
		newX := i + offset[0]
		newY := j + offset[1]
		if newX >= 0 && newX < rows && newY >=0 && newY < cols && grid[newX][newY] == byte('1') {
			dfs(grid,newX,newY)
		}
	}
}
func main(){
	fmt.Println(byte('1'))
grid := [][]byte{{'1','1','1','1','0'},{'0','0','0','0','0'}}
numIslands(grid)
}
