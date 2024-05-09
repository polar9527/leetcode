package go_case

/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 *
 * https://leetcode.com/problems/ones-and-zeroes/description/
 *
 * algorithms
 * Medium (47.56%)
 * Likes:    5364
 * Dislikes: 450
 * Total Accepted:    200.9K
 * Total Submissions: 422.5K
 * Testcase Example:  '["10","0001","111001","1","0"]\n5\n3'
 *
 * You are given an array of binary strings strs and two integers m and n.
 *
 * Return the size of the largest subset of strs such that there are at most m
 * 0's and n 1's in the subset.
 *
 * A set x is a subset of a set y if all elements of x are also elements of
 * y.
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
 * Output: 4
 * Explanation: The largest subset with at most 5 0's and 3 1's is {"10",
 * "0001", "1", "0"}, so the answer is 4.
 * Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
 * {"111001"} is an invalid subset because it contains 4 1's, greater than the
 * maximum of 3.
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["10","0","1"], m = 1, n = 1
 * Output: 2
 * Explanation: The largest subset is {"0", "1"}, so the answer is 2.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 600
 * 1 <= strs[i].length <= 100
 * strs[i] consists only of digits '0' and '1'.
 * 1 <= m, n <= 100
 *
 *
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 定义数组
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// 遍历
	for i := 0; i < len(strs); i++ {
		zeroNum, oneNum := 0, 0
		//计算0,1 个数
		//或者直接strings.Count(strs[i],"0")
		for _, v := range strs[i] {
			if v == '0' {
				zeroNum++
			}
		}
		oneNum = len(strs[i]) - zeroNum
		// 从后往前 遍历背包容量
		for j := m; j >= zeroNum; j-- {
			for k := n; k >= oneNum; k-- {
				// 推导公式
				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
			}
		}
		//fmt.Println(dp)
	}
	return dp[m][n]
}

// @lc code=end
