/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (44.24%)
 * Likes:    3125
 * Dislikes: 363
 * Total Accepted:    522.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	t := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	helper(0, "", digits, t, &ret)
	return ret
}

func helper(index int, prefix string, digits string, t []string, ret *[]string) {
	if index == len(digits) {
		*ret = append(*ret, prefix)
		return
	}
	word := t[digits[index]-'0']
	for _, r := range word {
		s := string(r)
		helper(index+1, prefix+s, digits, t, ret)
	}

}

func letterCombinations0503(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	dict := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var nums = make([]int, 0)
	for _, digit := range digits {
		num, err := strconv.Atoi(string(digit))
		if err != nil {
			return []string{}
		}
		if num == 1 {
			return []string{}
		}
		nums = append(nums, num)
	}
	ret := make([]string, 0)
	depth := len(nums)
	path := make([]string, 0)
	core(0, depth, nums, dict, path, &ret)
	return ret
}

func core(start, depth int, nums []int, dict, path []string, ret *[]string) {
	if start == depth {
		var result = make([]string, depth)
		copy(result, path)
		*ret = append(*ret, strings.Join(result, ""))
		return
	}

	for _, c := range dict[nums[start]] {
		path = append(path, string(c))
		core(start+1, depth, nums, dict, path, ret)
		path = path[:len(path)-1]
	}
}

// @lc code=end

func main() {
	nums := "23"
	ret := letterCombinations(nums)
	fmt.Println(ret)

}
