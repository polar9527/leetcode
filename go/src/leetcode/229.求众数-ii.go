/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 *
 * https://leetcode-cn.com/problems/majority-element-ii/description/
 *
 * algorithms
 * Medium (43.13%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    14.9K
 * Total Submissions: 34.5K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
 *
 * 说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: [3]
 *
 * 示例 2:
 *
 * 输入: [1,1,1,3,3,2,2,2]
 * 输出: [1,2]
 *
 */
package main

// @lc code=start
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return nil
	}
	candidate1 := nums[0]
	candidate2 := nums[0]
	count1, count2 := 0, 0

	for _, num := range nums {
		if candidate1 == num {
			count1++
			continue
		}
		if candidate2 == num {
			count2++
			continue
		}
		if count1 == 0 {
			candidate1 = num
			count1 = 1
			continue
		}
		if count2 == 0 {
			candidate2 = num
			count2 = 1
			continue
		}
		count1--
		count2--
	}

	if candidate1 != candidate2 {
		count1, count2 = 0, 0
		for _, num := range nums {
			if num == candidate1 {
				count1++
			}
		}
		for _, num := range nums {
			if num == candidate2 {
				count2++
			}
		}
		if count1 > len(nums)/3 {
			res = append(res, candidate1)
		}
		if count2 > len(nums)/3 {
			res = append(res, candidate2)
		}
	} else {
		count1 = 0
		for _, num := range nums {
			if num == candidate1 {
				count1++
			}
		}
		if count1 > len(nums)/3 {
			res = append(res, candidate1)
		}
	}

	return res
}

// @lc code=end
