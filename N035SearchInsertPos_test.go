package main

import (
	"fmt"
	"testing"
)

func TestN035SearchInsertPos(t *testing.T) {
	var n035 N035SearchInsertPos
	nums := []int{1, 3, 5, 6, 7, 9, 10, 11, 13, 15, 16, 17, 19, 20, 22, 24, 25, 26, 28, 29, 30, 33, 35, 36, 38, 40, 41, 43}
	fmt.Print("N035SearchInsertPos:\t")
	fmt.Print(n035.searchInsert(nums, len(nums), 0), " ")
	fmt.Print(n035.searchInsert(nums, len(nums), 33), " ")
	fmt.Print(n035.searchInsert(nums, len(nums), 32), " ")
	fmt.Print(n035.searchInsert(nums, len(nums), 50))
	fmt.Println()
}
