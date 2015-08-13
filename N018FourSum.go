package main

import "sort"

type N018FourSum struct {
}

func (this *N018FourSum) fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	size := len(nums)
	if size <= 3 {
		return result
	}

	var n0, n1, n2, n3, sum int
	for h := 0; h < size-3; {
		n0 = nums[h]
		for i := h + 1; i < size-2; {
			n1 = nums[i]
			for j, k := i+1, size-1; j < k; {
				n2 = nums[j]
				n3 = nums[k]
				sum = n0 + n1 + n2 + n3
				if sum > target {
					k--
				} else if sum < target {
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

					triplet := []int{n0, n1, n2, n3}
					result = append(result, triplet)
				}
			}
			i++
			for nums[i] == nums[i-1] && i < size-2 {
				i++
			}
		}
		h++
		for nums[h] == nums[h-1] && h < size-3 {
			h++
		}
	}
	return result
}
