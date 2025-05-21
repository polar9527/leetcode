/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode.cn/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (67.90%)
 * Likes:    2148
 * Dislikes: 0
 * Total Accepted:    678.2K
 * Total Submissions: 997.7K
 * Testcase Example:  '12'
 *
 * 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
 *
 * 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11
 * 不是。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 12
 * 输出：3
 * 解释：12 = 4 + 4 + 4
 *
 * 示例 2：
 *
 *
 * 输入：n = 13
 * 输出：2
 * 解释：13 = 4 + 9
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

// @lc code=start
// DP 一维
func numSquares(n int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// 背包大小为n
	// 物品为完全平方数
	// dp[j] 和为j 的完全平方数的最少数量
	// dp[j] = min(dp[j],dp[j-i*i]+1)
	dp := make([]int, n+1)
	// for i, _ := range dp {
	// 	dp[i] = math.MaxInt32
	// }
	dp[0] = 0
	for j := 1; j <= n; j++ { //背包
		dp[j] = 10e4 + 1
		for i := 1; i*i <= j; i++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

// dfs
func numSquares(n int) int {
	memo := make([][]int, int(math.Sqrt(float64(n)))+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i == 0 {
			if c == 0 {
				return 0
			}
			// 没有数字可选了，要保证 后面的min函数不取这个状态
			return math.MaxInt32
		}
		p := &memo[i][c]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() {
			*p = res
		}()
		if c < i*i {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-i*i)+1)
	}

	return dfs(int(math.Sqrt(float64(n))), n)
}

// DP 二维
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int, int(math.Sqrt(float64(n)))+1)
	for i := 0; i <= int(math.Sqrt(float64(n))); i++ {
		dp[i] = make([]int, n+1)
	}
	for j := range dp[0] {
		dp[0][j] = math.MaxInt32
	}
	dp[0][0] = 0
	// dp[i][0] = 0
	for i := 1; i*i <= n; i++ {
		for j := 0; j <= n; j++ {
			if j < i*i {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-i*i]+1)
			}
		}
	}
	return dp[int(math.Sqrt(float64(n)))][n]
}

// @lc code=end

