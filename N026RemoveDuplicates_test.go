package main

import (
	"fmt"
	"testing"
)

func TestN026RemoveDuplicates(t *testing.T) {
	var n026 N026RemoveDuplicates
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Print("N026RemoveDuplicates:\t")
	size := n026.removeDuplicates(nums)
	for i := 0; i < size; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}
