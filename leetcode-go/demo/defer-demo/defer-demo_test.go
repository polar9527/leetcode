package demo

import "testing"

func TestDeferDemo(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeferDemo(tt.args.flag)
			t.Log("test========")
		})
	}
}
