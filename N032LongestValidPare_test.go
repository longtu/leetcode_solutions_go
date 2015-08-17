package main

import (
	"fmt"
	"testing"
)

func TestN032LongestValidPare(t *testing.T) {
	var n032 N032LongestValidPare
	fmt.Print("N032LongestValidPare:\t")
	fmt.Print(n032.longestValidParentheses("()(())"), " ")
	fmt.Print(n032.longestValidParentheses("(()"), " ")
	fmt.Print(n032.longestValidParentheses(")(((((()())()()))()(()))("), " ")
	fmt.Println(n032.longestValidParentheses("((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((()"))

}
