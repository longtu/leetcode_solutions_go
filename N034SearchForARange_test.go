package main

import (
	"fmt"
	"testing"
)

func TestN034SearchForARange(t *testing.T) {
	var n034 N034SearchForARange
	nums := []int{5, 7, 7, 8, 8, 10}
	results := n034.searchRange(nums, len(nums), 8)
	fmt.Print("N034SearchForARange:\t")
	fmt.Println(results[0], results[1])
}
