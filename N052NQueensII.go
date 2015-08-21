package main

type N052NQueensII struct {
}

func (this *N052NQueensII) solveNQueensInternal(n int, curRow int, board [][]byte, pcount *int) {
	for curCol := 0; curCol < n; curCol++ {
		if board[curRow][curCol] != ' ' {
			newboard := make([][]byte, n)
			for i := 0; i < n; i++ {
				newboard[i] = make([]byte, n)
				copy(newboard[i], board[i])
			}
			newboard[curRow][curCol] = 'Q'
			if curRow == n-1 {
				// Found
				*pcount = *pcount + 1
			} else {
				//cout << "Checking (" << curRow << ", " << curCol << ")" << endl;
				// Remove '.' for this 'Q' and check next row
				for i := curRow; i < n; i++ {
					for j := 0; j < n; j++ {
						if newboard[i][j] == '.' {
							if i == curRow { // Remove '.' for current row
								newboard[i][j] = ' '
							}
							if j == curCol { // Remove '.' for curent column
								newboard[i][j] = ' '
							}
							if (i-curRow == j-curCol) || (i-curRow == -(j - curCol)) { // Remove '.' for 45 degree
								newboard[i][j] = ' '
							}
						}
					}
				}
				this.solveNQueensInternal(n, curRow+1, newboard, pcount)
			}
		}
	}
}

func (this *N052NQueensII) totalNQueens(n int) int {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	count := 0
	this.solveNQueensInternal(n, 0, board, &count)
	return count
}
