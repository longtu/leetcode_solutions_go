package main

import (
	"fmt"
	"testing"
)

func TestN041FirstMisPositive(t *testing.T) {
	var n041 N041FirstMisPositive
	nums1 := []int{1, 0}
	nums2 := []int{4, 2, 0, 1, 4}
	fmt.Print("N041FirstMisPositive:\t")
	fmt.Println(n041.firstMissingPositive(nums1, len(nums1)), n041.firstMissingPositive(nums2, len(nums2)))
}
