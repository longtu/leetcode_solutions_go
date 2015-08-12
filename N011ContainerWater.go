package main

type N011ContainerWater struct {
}

func (this *N011ContainerWater) maxArea(height []int, heightSize int) int {
	maxArea := 0
	area := 0
	value := 0
	i := 0
	j := heightSize - 1
	for i < j {
		if height[i] < height[j] {
			value = height[i]
		} else {
			value = height[j]
		}
		area = value * (j - i)
		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
