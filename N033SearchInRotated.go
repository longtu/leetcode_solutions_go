package main

type N033SearchInRotated struct {
}

func (this *N033SearchInRotated) binarySearch(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	if target < nums[start] || target > nums[end] {
		return -1
	}

	middle := (start + end) / 2
	for start <= end {
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
		middle = (start + end) / 2
	}
	return -1
}

func (this *N033SearchInRotated) searchInRange(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	MIN_SIZE := 10

	count := end - start + 1
	if count <= MIN_SIZE {
		for i := start; i <= end; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}
	index := -1
	newEnd := (start + end) / 2
	newStart := newEnd + 1
	if nums[start] < nums[newEnd] {
		index = this.binarySearch(nums, start, newEnd, target)
	} else {
		index = this.searchInRange(nums, start, newEnd, target)
	}

	if index >= 0 {
		return index
	}

	if nums[newStart] < nums[end] {
		index = this.searchInRange(nums, newStart, end, target)
	} else {
		index = this.searchInRange(nums, newStart, end, target)
	}
	return index
}

func (this *N033SearchInRotated) search(nums []int, numsSize int, target int) int {
	if numsSize == 0 {
		return -1
	}
	return this.searchInRange(nums, 0, numsSize-1, target)
}
