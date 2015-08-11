package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var n001 N001TwoSum
	nums := []int{2, 7, 11, 15, 12}
	results := n001.twoSum(nums, 9)
	fmt.Println("N001TwoSum: ", results)
	if len(results) != 2 {
		t.Errorf("Results length should be 2")
	} else if results[0] != 1 || results[1] != 2 {
		t.Errorf("Wrong results:", results)
	}
}
