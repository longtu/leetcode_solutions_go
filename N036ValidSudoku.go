package main

type N036ValidSudoku struct {
	flags [9]int
}

func (this *N036ValidSudoku) clearFlags() {
	for i := range this.flags {
		this.flags[i] = 0
	}
}

func (this *N036ValidSudoku) isRowValid(board [][]byte, row int, boardColSize int) bool {
	this.clearFlags()
	for i := 0; i < boardColSize; i++ {
		num := board[row][i]
		if num < '1' || num > '9' {
			continue
		}
		if this.flags[num-'1'] != 0 {
			return false
		}
		this.flags[num-'1'] = 1
	}
	return true
}

func (this *N036ValidSudoku) isColValid(board [][]byte, col int, boardRowSize int) bool {
	this.clearFlags()
	for i := 0; i < boardRowSize; i++ {
		num := board[i][col]
		if num < '1' || num > '9' {
			continue
		}
		if this.flags[num-'1'] != 0 {
			return false
		}
		this.flags[num-'1'] = 1
	}
	return true
}

func (this *N036ValidSudoku) isSubGridValid(board [][]byte, rowTopLeft int, colTopLeft int) bool {
	this.clearFlags()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			num := board[rowTopLeft+row][colTopLeft+col]
			if num < '1' || num > '9' {
				continue
			}
			if this.flags[num-'1'] != 0 {
				return false
			}
			this.flags[num-'1'] = 1
		}
	}
	return true
}

func (this *N036ValidSudoku) isValidSudoku(board [][]byte, boardRowSize int, boardColSize int) bool {
	for row := 0; row < boardRowSize; row++ {
		if !this.isRowValid(board, row, boardColSize) {
			return false
		}
	}
	for col := 0; col < boardColSize; col++ {
		if !this.isColValid(board, col, boardRowSize) {
			return false
		}
	}
	for row := 0; row < boardRowSize; row = row + 3 {
		for col := 0; col < boardColSize; col = col + 3 {
			if !this.isSubGridValid(board, row, col) {
				return false
			}
		}
	}
	return true
}
