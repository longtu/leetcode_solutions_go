package main

import (
	"fmt"
	"testing"
)

func TestN009PalindromeNumber(t *testing.T) {
	var n009 N009PalindromeNumber
	fmt.Print("N009PalindromeNumber:\t", n009.isPalindrome(0), " ")
	fmt.Print(n009.isPalindrome(123321), " ")
	fmt.Print(n009.isPalindrome(12321), " ")
	fmt.Print(n009.isPalindrome(1000110001))
	fmt.Println()
}
