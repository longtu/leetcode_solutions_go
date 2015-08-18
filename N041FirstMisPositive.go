package main

type N041FirstMisPositive struct {
}

func (this *N041FirstMisPositive) firstMissingPositive(nums []int, numsSize int) int {
	if numsSize == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	flags := make([]byte, numsSize)
	// Find count, min and max of all positives
	count := 0
	min := INT_MAX
	max := nums[0]
	for i := 0; i < numsSize; i++ {
		if nums[i] > 0 {
			count++
			flags[nums[i]-1] = 1
			if nums[i] < min {
				min = nums[i]
			}
			if nums[i] > max {
				max = nums[i]
			}
		}
	}

	if min > 1 {
		return 1
	}
	for i := 0; i < count; i++ {
		if flags[i] == 0 {
			return i + 1
		}
	}
	return count + 1
}
