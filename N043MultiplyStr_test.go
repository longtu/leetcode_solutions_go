package main

import (
	"fmt"
	"testing"
)

func TestN043MultiplyStr(t *testing.T) {
	var n043 N043MultiplyStr
	fmt.Print("N043MultiplyStr:\t")
	fmt.Println(n043.multiply("0", "0"), n043.multiply("9", "99"), n043.multiply("1200", "1.5"), n043.multiply("1200", "-01.5"))
}
