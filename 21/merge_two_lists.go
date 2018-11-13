package _21

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := new(ListNode)
	head := h
	for l1 != nil && l2 != nil {
		if l1 != nil && l1.Val < l2.Val {
			h.Next = l1
			h = h.Next
			l1 = l1.Next
		} else {
			h.Next = l2
			h = h.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		h.Next = l1
	} else {
		h.Next = l2
	}
	return head.Next
}
