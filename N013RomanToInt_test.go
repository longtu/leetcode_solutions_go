package main

import (
	"fmt"
	"testing"
)

func TestN013RomanToInt(t *testing.T) {
	var n013 N013RomanToInt
	fmt.Print("N013RomanToInt:\t\t")
	fmt.Println(n013.romanToInt("CMXCIX"), n013.romanToInt("MDCCCXLVI"))
}
