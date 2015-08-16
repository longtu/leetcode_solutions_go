package main

type N027RemoveElement struct {
}

func (this *N027RemoveElement) removeElement(nums []int, val int) int {
	numsSize := len(nums)
	i := 0
	j := 0
	removed := 0
	for j < numsSize {
		for nums[j] == val && j < numsSize {
			j++
			removed++
		}
		if j < numsSize {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return numsSize - removed
}
