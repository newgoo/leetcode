package _9

import (
	"testing"

	"github.com/astaxie/beego"
)

func Test_ttt(t *testing.T) {
	a := []int{1}
	h := new(ListNode)
	//p := new(ListNode)
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
		beego.Info(l.Val)
		l = l.Next
	}
	//for p != nil {
	//	beego.Info(p.Val)
	//	p = p.Next
	//}
}
