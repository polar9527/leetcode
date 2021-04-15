/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */
package leetcode

import "testing"

type testcase struct {
	input  []int
	expect int
}

func TestSingleNumber(t *testing.T) {
	testcases := []testcase{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for i := range testcases {
		ret := singleNumber(testcases[i].input)
		if ret != testcases[i].expect {
			t.Errorf("excpect: %v, got %v.", ret, testcases[i].expect)
		}
	}
}
