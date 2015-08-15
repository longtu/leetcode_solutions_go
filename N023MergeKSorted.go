package main

type N023MergeKSorted struct {
}

func (this *N023MergeKSorted) mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result, tail *ListNode
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

func (this *N023MergeKSorted) mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	size := len(lists)
	if size == 0 {
		return nil
	} else if size == 1 {
		return lists[0]
	}

	for size > 1 {
		// Divide and conque, merge two adjacent lists
		for i := 0; i < size-1; i += 2 {
			lists[i/2] = this.mergeTwoLists(lists[i], lists[i+1])
		}

		if size%2 == 0 {
			// We don't actually erase, but change the size only
			// lists.erase(lists.begin() + size / 2, lists.end());
			size /= 2
		} else {
			// Handle the last one when number is odd
			lists[size/2] = lists[size-1]
			// lists.erase(lists.begin() + size / 2 + 1, lists.end());
			size = size/2 + 1
		}
	}
	result = lists[0]
	return result
}
