package main

import (
	"fmt"
	"testing"
)

func TestN002AddTwoNumbers_addTwoNumbers(t *testing.T) {
	l0 := new(ListNode)
	l1 := new(ListNode)
	l2 := new(ListNode)
	l0.val = 3
	l0.next = nil
	l1.val = 4
	l1.next = l0
	l2.val = 2
	l2.next = l1

	l3 := new(ListNode)
	l4 := new(ListNode)
	l5 := new(ListNode)
	l3.val = 4
	l3.next = nil
	l4.val = 6
	l4.next = l3
	l5.val = 5
	l5.next = l4

	fmt.Print("N002AddTwoNumbers: ")
	l6 := N002AddTwoNumbers_addTwoNumbers(l2, l5)
	for l6 != nil {
		fmt.Printf("%d ", l6.val)
		l6 = l6.next
	}
	fmt.Println()
}
