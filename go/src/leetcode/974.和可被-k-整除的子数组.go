/*
 * @lc app=leetcode.cn id=974 lang=golang
 *
 * [974] 和可被 K 整除的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/description/
 *
 * algorithms
 * Medium (37.77%)
 * Likes:    136
 * Dislikes: 0
 * Total Accepted:    15.7K
 * Total Submissions: 35.6K
 * Testcase Example:  '[4,5,0,-2,-3,1]\n5'
 *
 * 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 *
 *
 *
 * 示例：
 *
 * 输入：A = [4,5,0,-2,-3,1], K = 5
 * 输出：7
 * 解释：
 * 有 7 个子数组满足其元素之和可被 K = 5 整除：
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2,
 * -3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 30000
 * -10000 <= A[i] <= 10000
 * 2 <= K <= 10000
 *
 *
 */
package main

// @lc code=start
func subarraysDivByK(A []int, K int) int {
	record := make(map[int]int)
	// 前缀和本身能mod K == 0 的情况
	record[0] = 1
	ans := 0
	sum := 0
	for _, n := range A {
		sum += n
		// 将 sum%K 为负数的情况调整为正数
		modulus := (sum%K + K) % K
		ans += record[modulus]
		record[modulus]++

	}
	return ans
}

// @lc code=end

func main() {

	A := []int{-1, 2, 9}
	K := 2
	subarraysDivByK(A, K)
}
