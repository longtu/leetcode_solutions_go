package main

import (
	"fmt"
	"testing"
)

func TestN007ReverseInteger_reverse(t *testing.T) {
	var n007 N007ReverseInteger
	fmt.Println("N007ReverseInteger:", n007.reverse(-123), n007.reverse(1534236469))
}
