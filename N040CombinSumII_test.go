package main

import (
	"fmt"
	"testing"
)

func TestN040CombinSumII(t *testing.T) {
	var n040 N040CombinSumII
	vec := []int{10, 1, 2, 7, 6, 1, 5}
	result := n040.combinationSum(vec, 8)
	fmt.Print("N040CombinSumII:\t")
	fmt.Println(result)
}
