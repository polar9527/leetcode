/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (44.97%)
 * Likes:    1704
 * Dislikes: 58
 * Total Accepted:    329.3K
 * Total Submissions: 731.5K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */
package main

// @lc code=start
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	permuteUniqueCore(nums, 0, &res)
	return res
}

func permuteUniqueCore(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		var premuteRes = make([]int, len(nums))
		copy(premuteRes, nums)
		*res = append(*res, premuteRes)
		return
	}
	m := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = true
		nums[start], nums[i] = nums[i], nums[start]
		permuteUniqueCore(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// @lc code=end
