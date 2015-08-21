package main

import (
	"fmt"
	"testing"
)

func TestN052NQueensII(t *testing.T) {
	var n052 N052NQueensII
	n := 8
	fmt.Print("N052NQueensII:\t\t", n052.totalNQueens(n), " results n=", n, "\n")
}
