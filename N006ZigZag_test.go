package main

import (
	"fmt"
	"testing"
)

func TestN006ZigZag_convert(t *testing.T) {
	var n006 N006ZigZag
	fmt.Println("N006ZigZag:", n006.convert("PAYPALISHIRING", 3), n006.convert("ABC", 2))
}
