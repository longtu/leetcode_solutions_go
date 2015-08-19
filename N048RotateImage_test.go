package main

import (
	"fmt"
	"testing"
)

func TestN048RotateImage(t *testing.T) {
	var n048 N048RotateImage
	row0 := []int{1, 2, 3}
	row1 := []int{4, 5, 6}
	matrix := [][]int{row0, row1}
	n048.rotate(&matrix)
	fmt.Println("N048RotateImage:")
	for i := range matrix {
		fmt.Print("\t\t\t")
		perm := matrix[i]
		for j := range perm {
			fmt.Print(perm[j], " ")
		}
		fmt.Println()
	}
}
