package go_case

/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 *
 * https://leetcode.com/problems/candy/description/
 *
 * algorithms
 * Hard (43.25%)
 * Likes:    7598
 * Dislikes: 631
 * Total Accepted:    497.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,0,2]'
 *
 * There are n children standing in a line. Each child is assigned a rating
 * value given in the integer array ratings.
 *
 * You are giving candies to these children subjected to the following
 * requirements:
 *
 *
 * Each child must have at least one candy.
 * Children with a higher rating get more candies than their neighbors.
 *
 *
 * Return the minimum number of candies you need to have to distribute the
 * candies to the children.
 *
 *
 * Example 1:
 *
 *
 * Input: ratings = [1,0,2]
 * Output: 5
 * Explanation: You can allocate to the first, second and third child with 2,
 * 1, 2 candies respectively.
 *
 *
 * Example 2:
 *
 *
 * Input: ratings = [1,2,2]
 * Output: 4
 * Explanation: You can allocate to the first, second and third child with 1,
 * 2, 1 candies respectively.
 * The third child gets 1 candy because it satisfies the above two
 * conditions.
 *
 *
 *
 * Constraints:
 *
 *
 * n == ratings.length
 * 1 <= n <= 2 * 10^4
 * 0 <= ratings[i] <= 2 * 10^4
 *
 *
 */

// @lc code=start
func candy(ratings []int) int {

	max := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	need := make([]int, len(ratings))
	sum := 0
	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = max(need[i-1], need[i]+1)
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}

// @lc code=end
