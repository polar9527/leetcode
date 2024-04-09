package go_case

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 *
 * https://leetcode.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (79.82%)
 * Likes:    10190
 * Dislikes: 385
 * Total Accepted:    526.2K
 * Total Submissions: 659.2K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * You are given a string s. We want to partition the string into as many parts
 * as possible so that each letter appears in at most one part.
 *
 * Note that the partition is done so that after concatenating all the parts in
 * order, the resultant string should be s.
 *
 * Return a list of integers representing the size of these parts.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it
 * splits s into less parts.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "eccbbbbdec"
 * Output: [10]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 500
 * s consists of lowercase English letters.
 *
 *
 */

// @lc code=start
func partitionLabels(s string) []int {
	lastOccurred := make(map[rune]int)
	var result []int
	for i, c := range s {
		lastOccurred[c] = i
	}
	start := 0
	max := 0
	for i, ch := range s {
		if lastOccurred[ch] > max {
			max = lastOccurred[ch]
		}
		if i == max {
			result = append(result, max-start+1)
			start = i + 1
		}
	}
	return result
}

// @lc code=end
