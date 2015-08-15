package main

import (
	"fmt"
	"testing"
)

func TestN022GenParentheses(t *testing.T) {
	var n022 N022GenParentheses
	result := n022.generateParenthesis(3) // 8 to test the complex situation
	fmt.Print("N022GenParentheses:\t", result, "\n")
}
