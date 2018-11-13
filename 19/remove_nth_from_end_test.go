package _19

import (
	"testing"
)

func Test_ttt(t *testing.T) {
	a := []int{1, 2, 35, 5}
	h := new(ListNode)
	p := h
	for _, one := range a {
		m := new(ListNode)
		m.Val = one
		h.Next = m
		h = h.Next
	}
	p = p.Next
	l := removeNthFromEnd(p, 1)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}
