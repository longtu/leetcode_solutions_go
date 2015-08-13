package main

import (
	"fmt"
	"testing"
)

func TestN008StringToInt(t *testing.T) {
	var n008 N008StringToInt
	fmt.Print("N008StringToInt: ", n008.myAtoi("1"), " ")
	fmt.Print(n008.myAtoi("0123"), " ")
	fmt.Print(n008.myAtoi("-1203"), " ")
	fmt.Print(n008.myAtoi("7e2"), " ")
	fmt.Print(n008.myAtoi("3E4"), " ")
	fmt.Print(n008.myAtoi("+11e530408314"), " ")
	fmt.Print(n008.myAtoi("+-2"), " ")
	fmt.Print(n008.myAtoi("    010"), " ")
	fmt.Print(n008.myAtoi("     +004500"), " ")
	fmt.Print(n008.myAtoi("  -0012a42"), " ")
	fmt.Print(n008.myAtoi("   +0 123"), " ")
	fmt.Print(n008.myAtoi("2147483648"), " ")
	fmt.Print(n008.myAtoi("   -115579378e25"), " ")
	fmt.Print(n008.myAtoi("9223372036854775809"))
	fmt.Println()
}
