package main

import (
	"fmt"
	"testing"
)

func TestN030SubstrWithConcat(t *testing.T) {
	var n030 N030SubstrWithConcat
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	results := n030.findSubstring(s, words)
	fmt.Print("N030SubstrWithConcat:\t")
	fmt.Println(results)
}
