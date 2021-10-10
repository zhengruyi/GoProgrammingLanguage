package DFS

//  从边缘往中间进行进行深度遍历，现将'O'-> 'C' -> 'O'
//  然后后续将'O'全部改成'X'
func solve(board [][]byte)  {
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			dfs(board,i,0)
		}
	}

	for i := 0; i < row; i++ {
		if board[i][col-1] == 'O' {
			dfs(board,i,col-1)
		}
	}

	for j := 0; j < col; j++ {
		if board[0][j] == 'O' {
			dfs(board,0,j)
		}
	}

	for j := 0; j < col; j++ {
		if board[row-1][j] == 'O' {
			dfs(board,row-1,j)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'C' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x,y int){
	board[x][y] = 'C'
	direction := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	row := len(board)
	col := len(board[0])
	for _, offset := range direction {
		newX := x + offset[0]
		newY := y + offset[1]
		if newX >= 0 && newX < row && newY >= 0 && newY < col && board[newX][newY] == 'O'{
			dfs(board,newX,newY)
		}
	}
}