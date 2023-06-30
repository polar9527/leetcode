package offer

import (
	"reflect"
	"testing"

	type24 "github.com/polar9527/leetcode/leetcode-go/offer/24.reverse-list/type"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *type24.ListNode
	}
	tests := []struct {
		name string
		args args
		want *type24.ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
