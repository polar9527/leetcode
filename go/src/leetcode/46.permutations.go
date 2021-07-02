/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (61.49%)
 * Likes:    3413
 * Dislikes: 98
 * Total Accepted:    557.7K
 * Total Submissions: 905.9K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */
package leetcode

import (
	"fmt"
)

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	premuteCore(nums, 0, &res)
	return res
}

func premuteCore(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		var premuteRes = make([]int, len(nums))
		copy(premuteRes, nums)
		*res = append(*res, premuteRes)
		return
	}

	for i := start; i < len(nums); i++ {

		nums[start], nums[i] = nums[i], nums[start]
		premuteCore(nums, start+1, res)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

// @lc code=end

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}
