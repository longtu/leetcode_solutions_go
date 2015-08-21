package main

import (
	"fmt"
	"testing"
)

func TestN051NQueens(t *testing.T) {
	var n051 N051NQueens
	n := 8
	results := n051.solveNQueens(n)
	fmt.Print("N051NQueens:\t\t", len(results), " results n=", n, "\n")
	//	for i, size := 0, len(results); i < size; i++ {
	//		result := results[i]
	//		fmt.Print("\t\t\t[\n")
	//		for j, size2 := 0, len(result); j < size2; j++ {
	//			fmt.Print("\t\t\t  ", result[j], "\n")
	//		}
	//		fmt.Println("\t\t\t]")
	//	}
}
