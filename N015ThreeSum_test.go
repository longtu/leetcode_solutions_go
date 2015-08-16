package main

import (
	"fmt"
	"testing"
)

func TestN015ThreeSum(t *testing.T) {
	var n015 N015ThreeSum
	nums := []int{-1, 0, 1, 2, -1, -1, -4}
	result := n015.threeSum(nums)
	fmt.Print("N0153Sum:\t\t")
	fmt.Println(result)
}
