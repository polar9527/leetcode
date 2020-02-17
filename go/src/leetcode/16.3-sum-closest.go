/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.72%)
 * Likes:    1621
 * Dislikes: 117
 * Total Accepted:    421.1K
 * Total Submissions: 921.2K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

// @lc code=start

func threeSumClosest(nums []int, target int) int {
	curDistance := math.MaxUint32
	ret := math.MaxUint32
	length := len(nums)
	if length < 3 {
		return -1
	}

	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < length-2; i++ {

		fi := i + 1
		se := length - 1

		for fi < se {
			sum := nums[i] + nums[fi] + nums[se]
			if sum == target {
				return sum
			} else if sum > target {
				newDistance := abs(sum - target)
				if newDistance < curDistance {
					curDistance = newDistance
					ret = sum
				} else {
					for fi < se && nums[se-1] == nums[se] {
						se--
					}
					se--
				}
			} else {
				newDistance := abs(sum - target)
				if newDistance < curDistance {
					curDistance = newDistance
					ret = sum
				} else {
					for fi < se && nums[fi+1] == nums[fi] {
						fi++
					}
					fi++
				}
			}
		}

	}

	return ret
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

// @lc code=end

func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, 0, 0}
	nums := []int{-1, 2, 1, -4}
	tar := 1
	ret := threeSumClosest(nums, tar)
	fmt.Println(ret)
}
