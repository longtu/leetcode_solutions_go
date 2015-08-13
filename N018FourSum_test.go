package main

import (
	"fmt"
	"testing"
)

func TestN018FourSum(t *testing.T) {
	nums := []int{-1, 0, 0, 1, 2, -2}
	var n018 N018FourSum
	fmt.Println("N018FourSum:", n018.fourSum(nums, 0))
}
