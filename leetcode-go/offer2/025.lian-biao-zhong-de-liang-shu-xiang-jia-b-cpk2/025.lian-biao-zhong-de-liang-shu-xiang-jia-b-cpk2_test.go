package offer2

import (
	"testing"

	"github.com/polar9527/leetcode/leetcode-go/offer2/listnode"
)

var l1, l2, l3 *listnode.ListNode

func init() {
	a := listnode.ListNode{
		Val: 7,
	}

	b := listnode.ListNode{
		Val: 2,
	}

	c := listnode.ListNode{
		Val: 4,
	}

	d := listnode.ListNode{
		Val: 3,
	}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	l1 = &a

	e := listnode.ListNode{
		Val: 5,
	}

	f := listnode.ListNode{
		Val: 6,
	}

	g := listnode.ListNode{
		Val: 4,
	}
	e.Next = &f
	f.Next = &g
	l2 = &e

	h := listnode.ListNode{
		Val: 7,
	}

	i := listnode.ListNode{
		Val: 8,
	}

	j := listnode.ListNode{
		Val: 0,
	}

	k := listnode.ListNode{
		Val: 7,
	}
	h.Next = &i
	i.Next = &j
	j.Next = &k
	l3 = &h
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *listnode.ListNode
		l2 *listnode.ListNode
	}
	tests := []struct {
		name string
		args args
		want *listnode.ListNode
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
