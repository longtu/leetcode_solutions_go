package main

import ()

func N001TwoSum_twoSum(nums []int, target int) []int {
	var results []int
	var hash map[int]int
	hash = make(map[int]int)
	index := 0
	// Use only one loop to build a hash and check if (target - value) exists in the hash
	for _, num := range nums {
		value, ok := hash[target-num]
		if ok && index != value {
			// Make sure the order is correct
			if index < value {
				results = append(results, index+1, value+1)
			} else {
				results = append(results, value+1, index+1)
			}
			break
		}
		hash[num] = index
		index++
	}

	return results
}
