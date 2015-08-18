package main

type N045JumpGameII struct {
}

func (this *N045JumpGameII) jump(nums []int, numsSize int) int {
	if numsSize <= 1 {
		return 0
	}
	jumpsCount := 1
	end := 0
	maxEnd := 0
	for end < numsSize-1 {
		// Find possible range in this step
		for i := maxEnd; i >= end; i-- {
			temp := i + nums[i]
			if temp > maxEnd {
				maxEnd = temp
			}
			if maxEnd >= numsSize-1 {
				return jumpsCount
			}
		}
		end++
		jumpsCount++
	}
	return jumpsCount
}
