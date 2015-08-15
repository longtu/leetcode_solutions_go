package main

type N025ReverseNodesKGrp struct {
}

// head is the first node of a group, and it will be the end of a group
func (this *N025ReverseNodesKGrp) reverseKGroupInternal(headOfGroup *ListNode, k int, preOfGroup *ListNode, pnewHead **ListNode, pnewEnd **ListNode) bool {
	if headOfGroup == nil || k <= 1 {
		return false
	}

	var cur, next *ListNode
	cur = headOfGroup
	next = cur.next
	*pnewHead = cur
	if next == nil {
		return false
	}
	*pnewEnd = cur
	(*pnewEnd).next = nil
	i := 1
	cur = next
	next = next.next
	cur.next = *pnewHead
	*pnewHead = cur
	i++
	for i < k && next != nil {
		cur = next
		next = next.next
		cur.next = *pnewHead
		*pnewHead = cur
		i++
	}

	if i == k {
		// Success
		(*pnewEnd).next = next
		if preOfGroup != nil {
			preOfGroup.next = *pnewHead
		}
		return true
	}

	// The node count is less than k in this group. We should revert it back.
	headOfGroup = *pnewHead
	cur = headOfGroup
	next = cur.next
	*pnewHead = cur
	*pnewEnd = cur
	(*pnewEnd).next = nil
	cur = next
	next = next.next
	cur.next = *pnewHead
	*pnewHead = cur
	if next != nil {
		cur = next
		next = next.next
		cur.next = *pnewHead
		*pnewHead = cur
	}
	if preOfGroup != nil {
		preOfGroup.next = *pnewHead
	}
	return false
}

func (this *N025ReverseNodesKGrp) reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	var pre, newHead, newEnd *ListNode
	successful := this.reverseKGroupInternal(head, k, pre, &newHead, &newEnd)
	newHeadOfWholeList := newHead
	for successful {
		head = newEnd.next
		pre = newEnd
		successful = this.reverseKGroupInternal(head, k, pre, &newHead, &newEnd)
	}
	return newHeadOfWholeList
}
