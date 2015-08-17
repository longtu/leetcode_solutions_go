package main

import (
	"fmt"
	"testing"
)

func TestN029DivideTwoInt(t *testing.T) {
	var n029 N029DivideTwoInt

	fmt.Print("N029DivideTwoInt:\t")
	fmt.Print(n029.divide(13, 3), " ")
	fmt.Print(n029.divide(2147483647, 3), " ")
	fmt.Print(n029.divide(-2147483648, -1), " ")
	fmt.Print(n029.divide(-1010369383, -2147483648), " ")
	fmt.Print(n029.divide(1, -1))
	fmt.Println()
}
