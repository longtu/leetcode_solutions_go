package main

import (
	"fmt"
	"testing"
)

func TestN012IntToRoman(t *testing.T) {
	var n012 N012IntToRoman
	fmt.Print("N012IntToRoman:\t\t")
	fmt.Println(n012.intToRoman(999), n012.intToRoman(1846))

}
