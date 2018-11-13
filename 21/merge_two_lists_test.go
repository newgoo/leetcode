package _21

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	a := []int{1, 2, 5, 35}
	b := []int{1, 2, 4, 39}
	aList := sliceToList(a)
	bList := sliceToList(b)
	h := mergeTwoLists(aList, bList)
	for h != nil {
		println(h.Val)
		h = h.Next
	}
}

func sliceToList(s []int) *ListNode {
	head := new(ListNode)
	h := head
	for _, one := range s {
		m := new(ListNode)
		m.Val = one
		head.Next = m
		head = head.Next
	}
	return h.Next
}
