package main

import (
	"fmt"
	"testing"
)

func TestN023MergeKSorted(t *testing.T) {
	var n023 N023MergeKSorted
	n0 := new(ListNode)
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n5 := new(ListNode)
	n6 := new(ListNode)
	n7 := new(ListNode)
	n8 := new(ListNode)
	n0.val = 0
	n1.val = 1
	n2.val = 2
	n3.val = 3
	n4.val = 4
	n5.val = 5
	n6.val = 6
	n7.val = 7
	n8.val = 8
	n0.next = n1
	n1.next = n4
	n4.next = n5

	n2.next = n3
	n3.next = n7

	n6.next = n8
	vec := []*ListNode{n0, n2, n6}
	result := n023.mergeKLists(vec)
	fmt.Print("N023MergeKSorted:\t")
	for result != nil {
		fmt.Print(result.val, " ")
		result = result.next
	}
	fmt.Println()
}
