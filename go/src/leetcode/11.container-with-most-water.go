/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (48.26%)
 * Likes:    4645
 * Dislikes: 506
 * Total Accepted:    527.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 *
 *
 *
 *
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49.
 *
 *
 *
 * Example:
 *
 *
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 */

// @lc code=start
package leetcode

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	lenght := len(height)
	if lenght <= 1 {
		return 0
	}
	maxArea := 0.0
	head := 0
	tail := lenght - 1

	for head < tail {
		temp := float64(tail-head) * math.Min(float64(height[head]), float64(height[tail]))
		if temp > maxArea {
			maxArea = temp
		}
		if height[head] > height[tail] {
			tail--
		} else {
			head++
		}
	}

	return int(maxArea)
}

func maxAreaB(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	lenght := len(height)

	maxArea := 0.0
	temp := 0.0
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ {
			temp = float64(j-i) * math.Min(float64(height[i]), float64(height[j]))
			if temp > maxArea {
				maxArea = temp
			}
		}
	}
	return int(maxArea)
}

// @lc code=end

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))

}
