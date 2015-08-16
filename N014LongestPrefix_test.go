package main

import (
	"fmt"
	"testing"
)

func TestN014LongestPrefix(t *testing.T) {
	var n014 N014LongestPrefix
	strs := []string{"abcdefg", "abcdasdfg", "abcdwgfsfs"}
	fmt.Print("N014LongestPrefix:\t")
	fmt.Println(n014.longestCommonPrefix(strs))
}
