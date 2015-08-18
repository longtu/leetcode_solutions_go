package main

import (
	"sort"
)

type N040CombinSumII struct {
}

// Comparing to V1, we added 'begin'. Because the candidates have been sorted, and we scan ascendingly,
// we need not try smaller number again. The combination is sorted and unique. We don't need set any more.
func (this *N040CombinSumII) findCombin(candidates []int, target int, combin *[]int, result *[][]int, begin int) {
	if target < candidates[0] {
		return
	}
	size := len(candidates)
	for i := begin; i < size; i++ {
		value := candidates[i]
		// Skip same value in the same layer
		if i > begin && value == candidates[i-1] {
			continue
		}
		*combin = append(*combin, value) // Push
		if value == target {
			tmp := make([]int, len(*combin))
			copy(tmp, *combin)
			*result = append(*result, tmp)       //Need not sort. Already sorted and unique.
			*combin = (*combin)[:len(*combin)-1] // Pop
			return
		}
		this.findCombin(candidates, target-value, combin, result, i+1)
		*combin = (*combin)[:len(*combin)-1] // Pop
	}
}

func (this *N040CombinSumII) combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combin := make([]int, 0)
	sort.Ints(candidates)
	this.findCombin(candidates, target, &combin, &result, 0)
	return result
}
