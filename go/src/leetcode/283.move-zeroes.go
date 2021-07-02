/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (53.90%)
 * Total Accepted:    440.3K
 * Total Submissions: 816.7K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 *
 */

package leetcode

func moveZeroes(nums []int) {
	lenght := len(nums)
	for tail, head := 0, 1; head < lenght; head++ {
		if nums[head] != 0 && nums[tail] == 0 {
			nums[tail], nums[head] = nums[head], nums[tail]
		}
		if nums[tail] != 0 {
			tail++
		}
	}
}
