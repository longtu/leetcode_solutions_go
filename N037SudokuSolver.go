package main

import "fmt"

const (
	SudokuStatusInvalid   = 0
	SudokuStatusNotSolved = 1
	SudokuStatusSolved    = 2
)

type N037SudokuSolver struct {
	settemplate map[byte]int
}

func (this *N037SudokuSolver) buildMap(board [][]byte, candimap map[int]map[byte]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				candimap[i*10+j] = make(map[byte]int)
				for k := byte('1'); k <= byte('9'); k++ {
					candimap[i*10+j][k] = 1
				}
			}
		}
	}
}

func (this *N037SudokuSolver) removeCandidate(board [][]byte, candimap map[int]map[byte]int, x int, y int, c byte) {
	if board[x][y] != '.' {
		return
	}

	myset := candimap[x*10+y]
	_, ok := myset[c]
	if ok {
		delete(myset, c)
		if len(myset) == 1 {
			for k := range myset {
				board[x][y] = k
			}
			delete(candimap, x*10+y)
			this.removeChar(board, candimap, x, y)
		}
	}
}

func (this *N037SudokuSolver) removeChar(board [][]byte, candimap map[int]map[byte]int, x int, y int) {
	c := board[x][y]
	if c == '.' {
		return
	}

	// Row
	for i := 0; i < 9; i++ {
		this.removeCandidate(board, candimap, x, i, c)
	}
	// Column
	for i := 0; i < 9; i++ {
		this.removeCandidate(board, candimap, i, y, c)
	}
	// Subboard
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			this.removeCandidate(board, candimap, x/3*3+i, y/3*3+j, c)
		}
	}
}

func (this *N037SudokuSolver) removeCharFromCandidates(board [][]byte, candimap map[int]map[byte]int) {
	// Remove existing numbers from the candidates
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			this.removeChar(board, candimap, i, j)
		}
	}
}

func (this *N037SudokuSolver) checkSudokuRow(board [][]byte, x int) int {
	check := [9]int{0}
	row := board[x]
	hasDot := false
	for i := 0; i < 9; i++ {
		c := row[i]
		if c == '.' {
			hasDot = true
		} else {
			if check[c-'1'] == 1 {
				return SudokuStatusInvalid
			} else {
				check[c-'1'] = 1
			}
		}
	}
	if hasDot {
		return SudokuStatusNotSolved
	}
	return SudokuStatusSolved
}

func (this *N037SudokuSolver) checkSudokuColumn(board [][]byte, x int) int {
	check := [9]int{0}
	hasDot := false
	for i := 0; i < 9; i++ {
		c := board[i][x]
		if c == '.' {
			hasDot = true
		} else {
			if check[c-'1'] == 1 {
				return SudokuStatusInvalid
			} else {
				check[c-'1'] = 1
			}
		}
	}
	if hasDot {
		return SudokuStatusNotSolved
	}
	return SudokuStatusSolved
}

func (this *N037SudokuSolver) checkSudokuSubboard(board [][]byte, x int, y int) int {
	check := [9]int{0}
	hasDot := false
	for i := x; i < 3; i++ {
		for j := y; j < 3; j++ {
			c := board[i][j]
			if c == '.' {
				hasDot = true
			} else {
				if check[c-'1'] == 1 {
					return SudokuStatusInvalid
				} else {
					check[c-'1'] = 1
				}
			}
		}
	}
	if hasDot {
		return SudokuStatusNotSolved
	}
	return SudokuStatusSolved
}

func (this *N037SudokuSolver) checkSudoku(board [][]byte) int {
	var status int
	allRowsSolved := true
	allColsSolved := true
	allSubsSolved := true
	for i := 0; i < 9; i++ {
		status = this.checkSudokuRow(board, i)
		allRowsSolved = allRowsSolved && (status == SudokuStatusSolved)
		if status == SudokuStatusInvalid {
			return SudokuStatusInvalid
		}
	}
	for i := 0; i < 9; i++ {
		status = this.checkSudokuColumn(board, i)
		allColsSolved = allColsSolved && (status == SudokuStatusSolved)
		if status == SudokuStatusInvalid {
			return SudokuStatusInvalid
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			status = this.checkSudokuSubboard(board, i, j)
			allSubsSolved = allSubsSolved && (status == SudokuStatusSolved)
			if status == SudokuStatusInvalid {
				return SudokuStatusInvalid
			}
		}
	}
	if allRowsSolved && allColsSolved && allSubsSolved {
		return SudokuStatusSolved
	}
	return SudokuStatusNotSolved
}

func (this *N037SudokuSolver) PrintWithMap(board [][]byte, candimap map[int]map[byte]int) {
	for i := 0; i < 9; i++ {
		row := board[i]
		for j := 0; j < 9; j++ {
			if row[j] != '.' {
				fmt.Print(string(row[j]))
			} else {
				myset := candimap[i*10+j]
				for k := range myset {
					fmt.Print(string(k))
				}
			}
			fmt.Print("\t")
		}
		fmt.Println()
	}
	fmt.Println("candimap size:", len(candimap))
}

func (this *N037SudokuSolver) Print(board [][]byte) {
	for i := 0; i < 9; i++ {
		row := board[i]
		for j := 0; j < 9; j++ {
			fmt.Print(string(row[j]), "\t")
		}
		fmt.Println()
	}
}

func (this *N037SudokuSolver) solveSudoku(board [][]byte) {
	candimap := make(map[int]map[byte]int)
	this.buildMap(board, candimap)
	if len(candimap) == 0 {
		return
	}
	this.removeCharFromCandidates(board, candimap)
	//print(board, candimap);
	status := this.checkSudoku(board)
	//cout << "Status: " << status << endl;
	if status == SudokuStatusSolved {
		return
	}

	// Try each candidate
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				theset := candimap[i*10+j]
				//fmt.Println(theset)
				for k := range theset {
					// vector<vector<char>> tryboard = board;
					tryboard := make([][]byte, 9)
					for p := 0; p < 9; p++ {
						tryboard[p] = make([]byte, 9)
						copy(tryboard[p], board[p])
					}
					// CandidateMap trycandimap = candimap;
					trycandimap := make(map[int]map[byte]int)
					for k2 := range candimap {
						trycandimap[k2] = make(map[byte]int)
						for k3 := range candimap[k2] {
							trycandimap[k2][k3] = candimap[k2][k3]
						}
					}
					tryboard[i][j] = k
					delete(trycandimap, i*10+j)
					this.removeChar(tryboard, trycandimap, i, j)
					//print(tryboard, trycandimap);
					status2 := this.checkSudoku(tryboard)
					//cout << "Status: " << status2 << endl;
					if status2 == SudokuStatusSolved {
						for i := 0; i < 9; i++ {
							copy(board[i], tryboard[i])
						}
						return
					}
				}
			}
		}
	}

}
