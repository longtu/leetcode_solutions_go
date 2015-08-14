package main

import (
	"fmt"
	"testing"
)

func TestN016ThreeSumClosest(t *testing.T) {
	var n016 N016ThreeSumClosest
	nums := []int{-1, 0, 1, 2, -1, -1, -4}
	fmt.Println("N016ThreeSumClosest:\t", n016.threeSumClosest(nums, 1))
}
