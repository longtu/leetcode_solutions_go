package main

import (
	"fmt"
	"testing"
)

func TestN042TrapRainWater(t *testing.T) {
	var n042 N042TrapRainWater
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Print("N042TrapRainWater:\t")
	fmt.Println(n042.trap(nums, len(nums)))
}
