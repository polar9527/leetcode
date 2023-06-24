/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
package leetcode

import (
	"testing"
)

type testcase struct {
	input  []int
	expect map[int]bool
}

func TestSingleNumber(t *testing.T) {
	testcases := []testcase{
		// {input: []int{1, 2, 1, 3, 2, 5}, expect: map[int]bool{3: true, 5: true}},
		{input: []int{1, 1, 0, -2147483648}, expect: map[int]bool{0: true, -2147483648: true}},
	}

	for itc := range testcases {
		ret := singleNumber(testcases[itc].input)
		for i := range ret {
			if _, ok := testcases[itc].expect[ret[i]]; !ok {
				t.Errorf("excpect: %v, got %v.", ret, testcases[i].expect)
			}
		}
		// if !reflect.DeepEqual(ret, testcases[i].expect) {

		// }
	}
	a := []int{}
	b := []int(nil)
	t.Logf("[]int{} != []int(nil), %v, %v", a, b)
}

// @lc code=end
