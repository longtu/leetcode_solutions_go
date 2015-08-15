package main

import (
	"fmt"
	"testing"
)

func TestN024SwapNodesInPairs(t *testing.T) {
	var n024 N024SwapNodesInPairs
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n5 := new(ListNode)
	n1.val = 1
	n2.val = 2
	n3.val = 3
	n4.val = 4
	n5.val = 5
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5

	result := n024.swapPairs(n1)
	fmt.Print("N024SwapNodesInPairs:\t")
	for result != nil {
		fmt.Print(result.val, " ")
		result = result.next
	}
	fmt.Println()
}
