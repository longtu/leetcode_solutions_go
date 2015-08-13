package main

import (
	"math"
	"sort"
)

type N016ThreeSumClosest struct {
}

func (this *N016ThreeSumClosest) threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)
	if size <= 2 {
		return 0
	}

	var n1, n2, n3, sum, diff int
	var minDiff, sumMinDiff int = INT_MAX, 0
	for i := 0; i < size-2; {
		n1 = nums[i]
		for j, k := i+1, size-1; j < k; {
			n2 = nums[j]
			n3 = nums[k]
			sum = n1 + n2 + n3
			if sum > target {
				diff = int(math.Abs(float64(sum - target)))
				if diff < minDiff {
					minDiff = diff
					sumMinDiff = sum
				}
				k--
				for nums[k] == nums[k+1] && k > j {
					k--
				}
			} else if sum < target {
				diff = int(math.Abs(float64(sum - target)))
				if diff < minDiff {
					minDiff = diff
					sumMinDiff = sum
				}
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else {
				return target
			}
		}
		i++
		for nums[i] == nums[i-1] && i < size-2 {
			i++
		}
	}
	return sumMinDiff
}
