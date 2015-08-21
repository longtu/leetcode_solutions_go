package main

type N051NQueens struct {
}

func (this *N051NQueens) solveNQueensInternal(n int, curRow int, board [][]byte, presults *[][]string) {
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
				result := make([]string, n)
				for i := 0; i < n; i++ {
					str := make([]byte, n)
					for j := 0; j < n; j++ {
						if newboard[i][j] == 'Q' {
							str = append(str, newboard[i][j])
						} else {
							str = append(str, '.')
						}
					}
					result[i] = string(str)
				}
				*presults = append(*presults, result)
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
				this.solveNQueensInternal(n, curRow+1, newboard, presults)
			}
		}
	}
}

func (this *N051NQueens) solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	results := make([][]string, 0)
	this.solveNQueensInternal(n, 0, board, &results)
	return results
}
