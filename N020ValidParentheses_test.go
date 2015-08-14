package main

import (
	"fmt"
	"testing"
)

func TestN020ValidParentheses(t *testing.T) {
	var n020 N020ValidParentheses
	fmt.Println("N020ValidParentheses:\t", n020.isValid("()[]{}"), n020.isValid("([)]"))
}
