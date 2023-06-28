package offer

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"case1", []int{-1, 5, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq1 := Constructor()
			cq1Ptr := &cq1

			if got := cq1Ptr.DeleteHead(); !reflect.DeepEqual(got, tt.want[0]) {
				t.Errorf("1 cq1Ptr.DeleteHead() = %v, want %v", got, tt.want[0])
			}

			cq1Ptr.AppendTail(5)
			cq1Ptr.AppendTail(2)
			if got := cq1Ptr.DeleteHead(); !reflect.DeepEqual(got, tt.want[1]) {
				t.Errorf("2 cq1Ptr.DeleteHead() = %v, want %v", got, tt.want[1])
			}
			if got := cq1Ptr.DeleteHead(); !reflect.DeepEqual(got, tt.want[2]) {
				t.Errorf("3 cq1Ptr.DeleteHead() = %v, want %v", got, tt.want[2])
			}
		})
	}
}
