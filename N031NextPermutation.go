package main

import "sort"

type N031NextPermutation struct {
}

func (this *N031NextPermutation) nextPermutation(nums []int) {
	maxStart := -1
	maxEnd := -1
	size := len(nums)
	if size <= 1 {
		return
	}

	// Find the farest increasing numbers combination. E.g. index 1, 3 in [3 4 8 5 1]
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] < nums[j] {
				if i > maxStart || (i == maxStart && j > maxEnd) {
					maxStart = i
					maxEnd = j
				}
			}
		}
	}

	if maxStart >= 0 {
		// Found. Move the bigger number forward. E.g. [3 5 4 8 1]
		temp := nums[maxEnd]
		for i := maxEnd; i > maxStart; i-- {
			nums[i] = nums[i-1]
		}
		nums[maxStart] = temp
		// Sort items after maxStart. E.g. [3 5 1 4 8]
		sort.Ints(nums[maxStart+1:])
	} else {
		// Not found. Numbers are descending. Just reverse nums.
		i := 0
		j := size - 1
		temp := 0
		for i < j {
			temp = nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j--
		}
	}
}
