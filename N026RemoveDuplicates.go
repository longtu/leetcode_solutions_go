package main

type N026RemoveDuplicates struct {
}

func (this *N026RemoveDuplicates) removeDuplicates(nums []int) int {
	numsSize := len(nums)
	if numsSize <= 1 {
		return numsSize
	}

	skipped := 0
	i := 1
	j := 1
	for j < numsSize {
		for nums[j] == nums[j-1] && j < numsSize {
			j++
			skipped++
		}
		if i != j {
			nums[i] = nums[j]
		}
		i++
		j++
	}
	return numsSize - skipped
}
