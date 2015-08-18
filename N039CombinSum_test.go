package main

import (
	"fmt"
	"testing"
)

func TestN039CombinSum(t *testing.T) {
	var n039 N039CombinSum
	vec := []int{3, 5, 7}
	result := n039.combinationSum(vec, 10)
	fmt.Print("N039CombinSum:\t\t")
	fmt.Println(result)
}
