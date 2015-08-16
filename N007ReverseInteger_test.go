package main

import (
	"fmt"
	"testing"
)

func TestN007ReverseInteger(t *testing.T) {
	var n007 N007ReverseInteger
	fmt.Print("N007ReverseInteger:\t")
	fmt.Println(n007.reverse(-123), n007.reverse(1534236469))
}
