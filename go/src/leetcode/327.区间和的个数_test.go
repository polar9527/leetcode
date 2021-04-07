package leetcode

import (
	"testing"
)

type testCase struct {
	arr    []int
	l      int
	u      int
	result int
}

func TestCountRangeSum(t *testing.T) {
	cases := []testCase{
		{[]int{-2, 5, -1}, -2, 2, 3},
	}

	for _, c := range cases {
		if ret := countRangeSum(c.arr, c.l, c.u); ret != c.result {
			t.Errorf("ret is %v, expect: %v", ret, c.result)
		}

	}
}
