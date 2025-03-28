package go_case

import "math"

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode.cn/problems/target-sum/description/
 *
 * algorithms
 * Medium (48.60%)
 * Likes:    2110
 * Dislikes: 0
 * Total Accepted:    542.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给你一个非负整数数组 nums 和一个整数 target 。
 *
 * 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
 *
 *
 * 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
 *
 *
 * 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,1,1], target = 3
 * 输出：5
 * 解释：一共有 5 种方法让最终目标和为 3 。
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], target = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 1000
 * 0 <= sum(nums[i]) <= 1000
 * -1000 <= target <= 1000
 *
 *
 */

// @lc code=start
// 1D array
// func findTargetSumWays(nums []int, target int) int {
// 	// a + b = sum
// 	// sum - b = a
// 	// a - b = target
// 	// sum - 2b = target
// 	// (sum - target)/2 = b
// 	abs := func(x int) int {
// 		return int(math.Abs(float64(x)))
// 	}

// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	if abs(target) > sum {
// 		return 0
// 	}
// 	if (sum-target)%2 == 1 {
// 		return 0
// 	}
// 	// 计算背包大小
// 	bag := (sum - target) / 2
// 	// 定义dp数组
// 	dp := make([]int, bag+1)
// 	// 初始化
// 	// 即装满背包为0的方法有一种，放0件物品。
// 	dp[0] = 1
// 	// 遍历顺序
// 	for i := 0; i < len(nums); i++ {
// 		for j := bag; j >= nums[i]; j-- {
// 			//推导公式
// 			dp[j] += dp[j-nums[i]]
// 			//fmt.Println(dp)
// 		}
// 	}
// 	return dp[bag]
// }

// 2D array
func findTargetSumWays(nums []int, target int) int {
	abs := func(t int) int {
		return int(math.Abs(float64(t)))
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if abs(target) > sum {
		return 0
	}

	if (sum-target)%2 == 1 {
		return 0
	}
	bagSize := (sum - target) / 2

	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, bagSize+1)
	}

	// init
	dp[0][0] = 1
	// first row
	if nums[0] <= bagSize {
		dp[0][nums[0]] = 1
	}
	// first column
	zeroNums := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroNums++
		}
		dp[i][0] = int(math.Pow(2.0, float64(zeroNums)))
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= bagSize; j++ {
			if j-nums[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][bagSize]
}

// @lc code=end
