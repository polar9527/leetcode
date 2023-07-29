package offer2

import "testing"

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name          string
		args          args
		wantMaxLength int
	}{
		// TODO: Add test cases.
		{"sample", args{[]int{0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxLength := findMaxLength(tt.args.nums); gotMaxLength != tt.wantMaxLength {
				t.Errorf("findMaxLength() = %v, want %v", gotMaxLength, tt.wantMaxLength)
			}
		})
	}
}
