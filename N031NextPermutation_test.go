package main

import (
	"fmt"
	"testing"
)

func TestN031NextPermutation(t *testing.T) {
	var n031 N031NextPermutation
	nums := []int{5, 1, 7, 3, 6, 4, 3}
	n031.nextPermutation(nums)
	fmt.Print("N031NextPermutation:\t")
	fmt.Println(nums)
}
