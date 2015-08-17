package main

import (
	"fmt"
	"testing"
)

func TestN036ValidSudoku(t *testing.T) {
	var n036 N036ValidSudoku
	board1 := [][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........")}
	fmt.Print("N036ValidSudoku:\t", n036.isValidSudoku(board1, 9, 9))
	fmt.Println()
}
