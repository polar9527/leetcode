/*
 * @lc app=leetcode.cn id=1371 lang=golang
 *
 * [1371] 每个元音包含偶数次的最长子字符串
 *
 * https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/description/
 *
 * algorithms
 * Medium (43.31%)
 * Likes:    129
 * Dislikes: 0
 * Total Accepted:    7.1K
 * Total Submissions: 12.6K
 * Testcase Example:  '"eleetminicoworoep"'
 *
 * 给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u'
 * ，在子字符串中都恰好出现了偶数次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "eleetminicoworoep"
 * 输出：13
 * 解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "leetcodeisgreat"
 * 输出：5
 * 解释：最长子字符串是 "leetc" ，其中包含 2 个 e 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "bcbcbc"
 * 输出：6
 * 解释：这个示例中，字符串 "bcbcbc" 本身就是最长的，因为所有的元音 a，e，i，o，u 都出现了 0 次。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 5 x 10^5
 * s 只包含小写英文字母。
 *
 *
 */
package main

import (
	"fmt"
)

// @lc code=start

func findTheLongestSubstring(s string) (ret int) {
	// status has 2**5 kinds
	pos := make([]int, 1<<5)
	for i := range pos {
		pos[i] = -1
	}
	status := 0
	result := 0
	// pos[0] 是指 aeiou 都为偶数:00000，这种奇偶模式的最早出现位置
	pos[0] = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}

		if pos[status] >= 0 {
			// 与当前位置相同奇偶性status最早出现的位置为j, 那么当前位置符合条件的子串的长度就是当前位置i减去当前位置相同奇偶性status最早出现的位置为j
			// 就是i-j+1, 而j==pos[status] => i - pos[status] + 1
			// 而已经遍历过之前的字符后储存的各种奇偶性status下计算出的符合条件的子串的长度为result
			// 取其中的最大值
			result = max(result, i+1-pos[status])
		} else {
			// 当此种奇偶性status第一次被发现，那么此时当前位置符合条件的子串的*长度*就是当前位置序号i加上1（这一点是显而易见的），例如i==0时，长度就是1(i+1)
			// 因为 i >= 0, 所以 i+1 >= 1, 当某种奇偶性status第一次被发现时，符合条件的子串的*长度*是一定>=1的，而这与以下这个特例不符
			// 但是有一个特例，就是status==00000,既所有的元音字符都出现偶数次的情况（包括所有元音都没出现的情况）
			// 这种奇偶性status第一次被发现，符合条件的子串是空串，其*长度*是0，所以有for循环开始前的pos[0] = 0
			//
			pos[status] = i + 1
		}
	}
	return result
}

// brute
// func findTheLongestSubstring(s string) int {
// 	maxLenght := 0
// 	for i := range s {
// 		countVowelMap := map[byte]int{
// 			'a' : 0,
// 			'e' : 0,
// 			'i' : 0,
// 			'o' : 0,
// 			'u' : 0,
// 		}
// 		// for j := range s[i:] {
// 		for j := i; j < len(s); j++ {
// 			if _, ok := countVowelMap[s[j]]; ok {
// 				countVowelMap[s[j]] += 1
// 			}
// 			isCorrect := true
// 			for _, v := range countVowelMap {
// 				if v%2 == 1 {
// 					isCorrect = false
// 					break
// 				}
// 			}
// 			if isCorrect {
// 				maxLenght = max(maxLenght, j-i+1)
// 			}
// 		}
// 	}
// 	return maxLenght
// }
//
func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

// @lc code=end

func main() {
	ret := findTheLongestSubstring("eleetminicoworoep")
	fmt.Println(ret)
}
