/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package leetcode

import "testing"

type testcase struct {
	input  [2]int
	expect int
}

func TestSuperEggDrop(t *testing.T) {

	// k=鸡蛋， n=楼层

	testcases := []testcase{
		{[2]int{1, 2}, 2},
		{[2]int{2, 6}, 3},
		{[2]int{100, 10000}, 14},
	}

	for i := range testcases {
		ret := superEggDrop(testcases[i].input[0], testcases[i].input[1])
		if ret != testcases[i].expect {
			t.Errorf("excpect: %v, got %v.", testcases[i].expect, ret)
		}
	}

}
