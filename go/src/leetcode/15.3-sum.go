/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.50%)
 * Likes:    5541
 * Dislikes: 670
 * Total Accepted:    770.8K
 * Total Submissions: 3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 */

// @lc code=start
package main

import (
	"fmt"
	"sort"
)

func threeSum1(nums []int) [][]int {
	var ret [][]int

	if len(nums) < 3 {
		return ret
	}
	m := make(map[[3]int]struct{})
	length := len(nums)
	for i := 0; i < length-2; i++ {
		for first := i + 1; first < length-1; first++ {
			for second := first + 1; second < length; second++ {
				sumTemp := nums[i] + nums[first] + nums[second]
				if sumTemp == 0 {
					s := []int{nums[i], nums[first], nums[second]}
					sort.Sort(sort.IntSlice(s))
					a := [3]int{s[0], s[1], s[2]}
					if _, ok := m[a]; !ok {
						m[a] = struct{}{}
						ret = append(ret, s)
					}
				}
			}
		}

	}
	return ret
}

func threeSum(nums []int) [][]int {
	var ret [][]int
	length := len(nums)
	if length < 3 {
		return ret
	}
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < length-2; i++ {
		fi := i + 1
		se := length - 1

		for fi < se {
			sum := nums[i] + nums[fi] + nums[se]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[fi], nums[se]})
				for fi < se && nums[fi+1] == nums[fi] {
					fi++
				}
				for fi < se && nums[se-1] == nums[se] {
					se--
				}
				fi++
				se--
			} else if sum > 0 {
				for fi < se && nums[se-1] == nums[se] {
					se--
				}
				se--
			} else {
				for fi < se && nums[fi+1] == nums[fi] {
					fi++
				}
				fi++
			}
		}
		if i < length-2 && nums[i] != nums[i+1] {
			continue
		}
		for i < length-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	// fmt.Println(nums)
	return ret
}

// @lc code=end

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0}
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	ret := threeSum(nums)
	fmt.Println(ret)
}
