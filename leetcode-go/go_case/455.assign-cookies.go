package go_case

import "sort"

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 *
 * https://leetcode.com/problems/assign-cookies/description/
 *
 * algorithms
 * Easy (52.66%)
 * Likes:    3742
 * Dislikes: 357
 * Total Accepted:    426.6K
 * Total Submissions: 810.2K
 * Testcase Example:  '[1,2,3]\n[1,1]'
 *
 * Assume you are an awesome parent and want to give your children some
 * cookies. But, you should give each child at most one cookie.
 *
 * Each child i has a greed factor g[i], which is the minimum size of a cookie
 * that the child will be content with; and each cookie j has a size s[j]. If
 * s[j] >= g[i], we can assign the cookie j to the child i, and the child i
 * will be content. Your goal is to maximize the number of your content
 * children and output the maximum number.
 *
 *
 * Example 1:
 *
 *
 * Input: g = [1,2,3], s = [1,1]
 * Output: 1
 * Explanation: You have 3 children and 2 cookies. The greed factors of 3
 * children are 1, 2, 3.
 * And even though you have 2 cookies, since their size is both 1, you could
 * only make the child whose greed factor is 1 content.
 * You need to output 1.
 *
 *
 * Example 2:
 *
 *
 * Input: g = [1,2], s = [1,2,3]
 * Output: 2
 * Explanation: You have 2 children and 3 cookies. The greed factors of 2
 * children are 1, 2.
 * You have 3 cookies and their sizes are big enough to gratify all of the
 * children,
 * You need to output 2.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= g.length <= 3 * 10^4
 * 0 <= s.length <= 3 * 10^4
 * 1 <= g[i], s[j] <= 2^31 - 1
 *
 *
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			count++
			i++
		}
	}
	return count
}

// @lc code=end
