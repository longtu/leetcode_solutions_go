package main

type N019RemoveNthNode struct {
}

func (this *N019RemoveNthNode) removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head

	for i := 0; i < n; i++ {
		p1 = p1.next
	}

	if p1 == nil {
		head = head.next
		return head
	}

	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	p1 = p2.next
	p2.next = p1.next
	return head
}
