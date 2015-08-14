package main

import (
	"fmt"
	"testing"
)

func TestN010RegEx(t *testing.T) {
	var n010 N010RegEx
	fmt.Print("N010RegEx:\t\t")
	fmt.Print(n010.isMatch("aa", "a"), " ")
	fmt.Print(n010.isMatch("aa", "aa"), " ")
	fmt.Print(n010.isMatch("aaa", "aa"), " ")
	fmt.Print(n010.isMatch("aaa", "ab*a"), " ")
	fmt.Print(n010.isMatch("aa", "a*"), " ")
	fmt.Print(n010.isMatch("aa", ".a"), " ")
	fmt.Print(n010.isMatch("ab", ".*"), " ")
	fmt.Print(n010.isMatch("aab", "c*a*b"), " ")
	fmt.Print(n010.isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*a*a*b"), " ")
	fmt.Print(n010.isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"))
	fmt.Println()
}
