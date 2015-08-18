package main

type N042TrapRainWater struct {
}

func (this *N042TrapRainWater) findIndexOfMax(nums []int, start int, end int) int {
	if end == start {
		return nums[start]
	}
	max := nums[start]
	index := start
	for i := start + 1; i <= end; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
}

// start and end are highest.
func (this *N042TrapRainWater) waterBetween(nums []int, start int, end int) int {
	if start+1 >= end {
		return 0
	}

	water := 0
	var height int
	if nums[start] > nums[end] {
		height = nums[end]
	} else {
		height = nums[start]
	}
	for i := start + 1; i < end; i++ {
		water += height - nums[i]
	}
	return water
}

func (this *N042TrapRainWater) trapInternal(nums []int, start int, end int, indexOfMax int, pwater *int) {
	if start+1 >= end {
		return
	}
	indexOfNextMax := 0
	if start == indexOfMax {
		indexOfNextMax = this.findIndexOfMax(nums, start+1, end)
		*pwater = *pwater + this.waterBetween(nums, start, indexOfNextMax)
		this.trapInternal(nums, indexOfNextMax, end, indexOfNextMax, pwater)
	} else {
		indexOfNextMax = this.findIndexOfMax(nums, start, end-1)
		*pwater = *pwater + this.waterBetween(nums, indexOfNextMax, end)
		this.trapInternal(nums, start, indexOfNextMax, indexOfNextMax, pwater)
	}
}

func (this *N042TrapRainWater) trap(height []int, heightSize int) int {
	if heightSize <= 1 {
		return 0
	}

	water := 0
	indexOfMax := this.findIndexOfMax(height, 0, heightSize-1)
	this.trapInternal(height, 0, indexOfMax, indexOfMax, &water)
	this.trapInternal(height, indexOfMax, heightSize-1, indexOfMax, &water)
	return water
}
