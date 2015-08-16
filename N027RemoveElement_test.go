package main

import (
	"fmt"
	"testing"
)

func TestN027RemoveElement(t *testing.T) {
	var n027 N027RemoveElement
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Print("N027RemoveElement:\t")
	size := n027.removeElement(nums, 2)
	for i := 0; i < size; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}
