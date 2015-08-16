package main

import (
	"fmt"
	"testing"
)

func TestN003LongestSubstring(t *testing.T) {
	var n003 N003LongestSubstring
	fmt.Print("N003LongestSubstring:\t")
	fmt.Println(n003.lengthOfLongestSubstring("abcabcbb"), n003.lengthOfLongestSubstring("eee"), n003.lengthOfLongestSubstring("dvdf"))
}
