package main

import (
	"fmt"
	"testing"
)

func TestN019RemoveNthNode(t *testing.T) {
	n5 := new(ListNode)
	n5.val = 5
	n4 := new(ListNode)
	n4.val = 4
	n3 := new(ListNode)
	n3.val = 3
	n2 := new(ListNode)
	n2.val = 2
	n1 := new(ListNode)
	n1.val = 1
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	var n019 N019RemoveNthNode

	result := n019.removeNthFromEnd(n1, 2)
	result = n019.removeNthFromEnd(result, 4)
	fmt.Print("N019RemoveNthNode: ")
	for result != nil {
		fmt.Print(result.val, " ")
		result = result.next
	}
	fmt.Println()
}
