package main

type N035SearchInsertPos struct {
}

func (this *N035SearchInsertPos) searchInternal(nums []int, start int, end int, target int) int {
	if target <= nums[start] {
		return start
	}
	if target > nums[end] {
		return end + 1
	}
	COUNT_TO_SEQUENTIAL_SEARCH := 8
	if end-start+1 <= COUNT_TO_SEQUENTIAL_SEARCH {
		for i := start; i <= end; i++ {
			if target <= nums[i] {
				return i
			}
		}
		return end + 1
	}

	middle := (start + end) / 2
	if nums[middle] == target {
		return middle
	} else if nums[middle] > target {
		return this.searchInternal(nums, start, middle-1, target)
	} else {
		return this.searchInternal(nums, middle+1, end, target)
	}
}

func (this *N035SearchInsertPos) searchInsert(nums []int, numsSize int, target int) int {
	if numsSize == 0 {
		return 0
	}

	return this.searchInternal(nums, 0, numsSize-1, target)
}
