package main

import "sort"

type N015ThreeSum struct {
}

func (this *N015ThreeSum) threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	size := len(nums)
	if size <= 2 {
		return result
	}

	var n1, n2, n3, sum int
	for i := 0; i < size-2; {
		n1 = nums[i]
		if n1 > 0 {
			break
		}
		for j, k := i+1, size-1; j < k; {
			n2 = nums[j]
			n3 = nums[k]
			sum = n1 + n2 + n3
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
				k--
				for nums[k] == nums[k+1] && k > j {
					k--
				}

				triplet := []int{n1, n2, n3}
				result = append(result, triplet)
			}
		}
		i++
		for nums[i] == nums[i-1] && i < size-2 {
			i++
		}
	}
	return result
}
