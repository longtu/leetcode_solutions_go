package main

import (
	"fmt"
	"testing"
)

func TestN046Permutations(t *testing.T) {
	var n046 N046Permutations
	nums := []int{1, 2, 3}
	results := n046.permute(nums)
	fmt.Print("N046Permutations:\t")
	for i := range results {
		perm := results[i]
		for j := range perm {
			fmt.Print(perm[j], " ")
		}
		fmt.Print("\t")
	}
	fmt.Println()
}
