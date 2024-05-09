package go_case

/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 *
 * https://leetcode.com/problems/combination-sum-iv/description/
 *
 * algorithms
 * Medium (54.15%)
 * Likes:    7303
 * Dislikes: 656
 * Total Accepted:    479.8K
 * Total Submissions: 886K
 * Testcase Example:  '[1,2,3]\n4'
 *
 * Given an array of distinct integers nums and a target integer target, return
 * the number of possible combinations that add up to target.
 *
 * The test cases are generated so that the answer can fit in a 32-bit
 * integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3], target = 4
 * Output: 7
 * Explanation:
 * The possible combination ways are:
 * (1, 1, 1, 1)
 * (1, 1, 2)
 * (1, 2, 1)
 * (1, 3)
 * (2, 1, 1)
 * (2, 2)
 * (3, 1)
 * Note that different sequences are counted as different combinations.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [9], target = 3
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * 1 <= nums[i] <= 1000
 * All the elements of nums are unique.
 * 1 <= target <= 1000
 *
 *
 *
 * Follow up: What if negative numbers are allowed in the given array? How does
 * it change the problem? What limitation we need to add to the question to
 * allow negative numbers?
 *
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	// 不能在这里初始化，因为dp[n-1][j-nums[0]] 的值还没有被计算出来
	// for j := nums[0]; j <= target; j++ {
	// 	dp[0][j] = dp[n-1][j-nums[0]]
	// }

	for j := 1; j <= target; j++ {
		if j >= nums[0] {
			dp[0][j] = dp[n-1][j-nums[0]]
		}
		for i := 1; i < n; i++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[n-1][j-nums[i]]
			}

		}

	}
	return dp[n-1][target]
}

// https://leetcode.cn/problems/combination-sum-iv/solutions/2624750/yi-ci-xing-dai-ni-gao-ding-suo-you-bei-b-8059/

// func combinationSum4(nums []int, target int) int {
// 	//定义dp数组
// 	dp := make([]int, target+1)
// 	// 初始化
// 	dp[0] = 1
// 	// 遍历顺序, 先遍历背包,再循环遍历物品
// 	for j:=0;j<=target;j++ {
// 		for i:=0 ;i < len(nums);i++ {
// 			if j >= nums[i] {
// 				dp[j] += dp[j-nums[i]]
// 			}
// 		}
// 	}
// 	return dp[target]
// }

// @lc code=end
