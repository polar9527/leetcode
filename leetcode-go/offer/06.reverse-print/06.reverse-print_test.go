package offer

import (
	"reflect"
	"testing"

	type06 "github.com/polar9527/leetcode/leetcode-go/offer/06.reverse-print/type"
)

var l1, l2 *type06.ListNode

func init() {
	a := type06.ListNode{
		Val: 1,
	}

	b := type06.ListNode{
		Val: 3,
	}

	c := type06.ListNode{
		Val: 2,
	}

	a.Next = &b
	b.Next = &c
	l1 = &a

	l2 = &type06.ListNode{}
	head := l2
	for i := 1; i <= 10000; i++ {
		head.Next = &type06.ListNode{}
		head = head.Next
	}
}

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *type06.ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"normal", args{head: l1}, []int{2, 3, 1}},
		{"low limit", args{head: nil}, []int{}},
		{"high limit", args{head: l2}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("case: %v, reversePrint() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
