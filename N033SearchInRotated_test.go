package main

import (
	"fmt"
	"testing"
)

func TestN033SearchInRotated(t *testing.T) {
	var n033 N033SearchInRotated
	nums := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 0, 1, 2, 3}
	fmt.Print("N033SearchInRotated:\t", n033.search(nums, len(nums), 0), " ")
	nums2 := []int{1, 3, 5}
	fmt.Println(n033.search(nums, len(nums2), 0))
}
