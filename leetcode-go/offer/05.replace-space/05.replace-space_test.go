package offer

import (
	"reflect"
	"strings"
	"testing"
)

var b = strings.Builder{}

func init() {
	for i := 0; i <= 10001; i++ {
		b.WriteString("x")
	}
}
func Test_replaceSpace(t *testing.T) {

	builder := strings.Builder{}

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{s: "We are happy."}, "We%20are%20happy."},
		{"low range limit", args{s: ""}, ""},
		{"high range limit", args{s: builder.String()}, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if gotResult := replaceSpace(tt.args.s); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("replaceSpaceRegexp() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}

func Test_replaceSpaceRegexp(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{s: "We are happy."}, "We%20are%20happy."},
		{"low range limit", args{s: ""}, ""},
		{"high range limit", args{s: b.String()}, b.String()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()
			if gotResult := replaceSpaceRegexp(tt.args.s); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("replaceSpaceRegexp() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}

func Test_replaceSpaceSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{s: "We are happy."}, "We%20are%20happy."},
		{"low range limit", args{s: ""}, ""},
		{"high range limit", args{s: b.String()}, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if gotResult := replaceSpaceSplit(tt.args.s); !reflect.DeepEqual(gotResult, tt.want) {
				t.Errorf("replaceSpaceRegexp() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
