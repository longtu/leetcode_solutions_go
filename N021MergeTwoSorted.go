package main

type N021MergeTwoSorted struct {
}

func (this *N021MergeTwoSorted) mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result, tail *ListNode = nil, nil
	if l1.val < l2.val {
		tail = l1
		result = l1
		l1 = l1.next
	} else {
		tail = l2
		result = l2
		l2 = l2.next
	}
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			tail.next = l1
			tail = l1
			l1 = l1.next
		} else {
			tail.next = l2
			tail = l2
			l2 = l2.next
		}
	}
	if l1 != nil {
		tail.next = l1
	}
	if l2 != nil {
		tail.next = l2
	}
	return result
}
