package _19

type ListNode struct {
	Val  int
	Next *ListNode
}

//0ms
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := make([]*ListNode, 0)
	var h = head
	for h != nil {
		l = append(l, h)
		h = h.Next
	}

	//只有一个元素
	if len(l)-n == 0 {
		return l[0].Next
	}

	l[len(l)-n-1].Next = l[len(l)-n].Next
	return head
}

//0ms
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	l := make([]*ListNode, 0)
	var h = head
	for h != nil {
		l = append(l, h)
		h = h.Next
	}

	//只有一个元素
	if len(l)-n == 0 {
		return l[0].Next
	}

	l[len(l)-n-1].Next = l[len(l)-n].Next
	return head
}
