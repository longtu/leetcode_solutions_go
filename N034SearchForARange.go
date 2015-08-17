package main

type N034SearchForARange struct {
}

// Return an array with 2 items
func (this *N034SearchForARange) searchInternal(nums []int, start int, end int, target int, results []int) bool {
	MIN_COUNT := 8
	results[0] = -1
	results[1] = -1
	if start > end || nums[start] > target || nums[end] < target {
		return false
	}

	if nums[start] == target && nums[end] == target {
		results[0] = start
		results[1] = end
		return true
	}

	if end-start+1 <= MIN_COUNT {
		for i := start; i <= end; i++ {
			if nums[i] == target {
				if results[0] == -1 {
					results[0] = i
				}
				results[1] = i
			}
		}
		return true
	}

	subresults1 := []int{-1, -1}
	subresults2 := []int{-1, -1}
	middle := (start + end) / 2
	found1 := this.searchInternal(nums, start, middle, target, subresults1)
	found2 := this.searchInternal(nums, middle, end, target, subresults2)
	if found1 {
		if found2 {
			results[0] = subresults1[0]
			results[1] = subresults2[1]
		} else {
			results[0] = subresults1[0]
			results[1] = subresults1[1]
		}
		return true
	} else {
		if found2 {
			results[0] = subresults2[0]
			results[1] = subresults2[1]
			return true
		} else {
			return false
		}
	}
}

/**
* Return an array of size *returnSize.
 */
func (this *N034SearchForARange) searchRange(nums []int, numsSize int, target int) []int {
	results := []int{-1, -1}
	if numsSize == 0 {
		return results
	}
	this.searchInternal(nums, 0, numsSize-1, target, results)
	return results
}
