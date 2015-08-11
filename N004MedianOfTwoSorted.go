package main

type N004MedianOfTwoSorted struct {
}

func (this *N004MedianOfTwoSorted) findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Merge two sorted arrays into one
	size1 := len(nums1)
	size2 := len(nums2)
	size3 := size1 + size2
	merged := make([]int, size3)
	for i1, i2, i3 := 0, 0, 0; i1 < size1 || i2 < size2; {
		if i1 == size1 {
			merged[i3] = nums2[i2]
			i2++
		} else if i2 == size2 {
			merged[i3] = nums1[i1]
			i1++
		} else {
			if nums1[i1] <= nums2[i2] {
				merged[i3] = nums1[i1]
				i1++
			} else {
				merged[i3] = nums2[i2]
				i2++
			}
		}
		i3++
	}

	var median float64 = 0.0
	// Median is the middle one if array size is odd, the average of middle two if array size is even
	if size3%2 == 1 {
		median = float64(merged[size3/2])
	} else {
		median = float64(merged[size3/2-1]+merged[size3/2]) / 2.0
	}
	return median
}
