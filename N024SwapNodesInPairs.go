package main

type N024SwapNodesInPairs struct {
}

func (this *N024SwapNodesInPairs) swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre, node1, node2, next *ListNode
	// Swap head
	node1 = head
	node2 = head.next
	if node2 == nil {
		return node1
	}
	next = node2.next
	node2.next = node1
	node1.next = next
	head = node2
	pre = node1

	// Swap other nodes
	for pre.next != nil && pre.next.next != nil {
		node1 = pre.next
		node2 = node1.next
		next = node2.next
		node2.next = node1
		node1.next = next
		pre.next = node2
		pre = node1
	}
	return head
}
