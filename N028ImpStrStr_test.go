package main

import (
	"fmt"
	"testing"
)

func TestN028ImpStrStr(t *testing.T) {
	var n028 N028ImpStrStr
	haystack := "abcdefgh ijkl mnopq"
	needle := "def"
	fmt.Print("N028ImpStrStr:\t\t")
	fmt.Println(n028.strStr(haystack, needle))
}
