package main

import (
	"fmt"
	"testing"
)

func TestN047PermutationsII(t *testing.T) {
	var n047 N047PermutationsII
	nums := []int{1, 1, 2}
	results := n047.permuteUnique(nums)
	fmt.Print("N047PermutationsII:\t")
	for i := range results {
		perm := results[i]
		for j := range perm {
			fmt.Print(perm[j], " ")
		}
		fmt.Print("\t")
	}
	fmt.Println()
}
