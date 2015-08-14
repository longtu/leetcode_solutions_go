package main

import (
	"fmt"
	"testing"
)

func TestN013RomanToInt(t *testing.T) {
	var n013 N013RomanToInt
	fmt.Println("N013RomanToInt:\t\t", n013.romanToInt("CMXCIX"), n013.romanToInt("MDCCCXLVI"))
}
