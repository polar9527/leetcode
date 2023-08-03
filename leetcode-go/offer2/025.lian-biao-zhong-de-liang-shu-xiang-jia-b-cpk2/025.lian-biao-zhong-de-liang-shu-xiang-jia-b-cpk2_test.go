package offer2

import (
	"testing"
)

var l1, l2, l3 *ListNode

func init() {
	a := ListNode{
		Val: 7,
	}

	b := ListNode{
		Val: 2,
	}

	c := ListNode{
		Val: 4,
	}

	d := ListNode{
		Val: 3,
	}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	l1 = &a

	e := ListNode{
		Val: 5,
	}

	f := ListNode{
		Val: 6,
	}

	g := ListNode{
		Val: 4,
	}
	e.Next = &f
	f.Next = &g
	l2 = &e

	h := ListNode{
		Val: 7,
	}

	i := ListNode{
		Val: 8,
	}

	j := ListNode{
		Val: 0,
	}

	k := ListNode{
		Val: 7,
	}
	h.Next = &i
	i.Next = &j
	j.Next = &k
	l3 = &h
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"case1", args{l1, l2}, l3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); got != nil {
				t.Logf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
