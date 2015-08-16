package main

import (
	"fmt"
	"testing"
)

func TestN004MedianOfTwoSorted(t *testing.T) {
	nums1 := []int{1, 4, 6, 9}
	nums2 := []int{2, 3, 5}
	nums3 := []int{}
	nums4 := []int{1}
	var n004 N004MedianOfTwoSorted
	fmt.Print("N004MedianOfTwoSorted:\t")
	fmt.Println(n004.findMedianSortedArrays(nums1, nums2), n004.findMedianSortedArrays(nums3, nums4))
}
