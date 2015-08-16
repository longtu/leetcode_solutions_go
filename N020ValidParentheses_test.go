package main

import (
	"fmt"
	"testing"
)

func TestN020ValidParentheses(t *testing.T) {
	var n020 N020ValidParentheses
	fmt.Print("N020ValidParentheses:\t")
	fmt.Println(n020.isValid("()[]{}"), n020.isValid("([)]"))
}
