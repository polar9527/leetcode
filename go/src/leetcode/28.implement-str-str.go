/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (39.68%)
 * Likes:    442
 * Dislikes: 0
 * Total Accepted:    170.7K
 * Total Submissions: 430.2K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */
package main

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext(needle)
	target := 0
	parttern := 0
	for target < len(haystack) {
		if haystack[target] == needle[parttern] {
			target++
			parttern++
		} else if parttern != 0 {
			parttern = next[parttern-1]
		} else {
			target++
		}
		if parttern == len(needle) {
			return target - len(needle)
		}
	}

	return -1
}

func getNext(parttern string) []int {
	next := make([]int, len(parttern))
	now := 0
	for i := 1; i < len(parttern); {

		if parttern[now] == parttern[i] {
			now++
			next[i] = now
			i++
		} else if now != 0 {
			now = next[now-1]
		} else {
			i++
			// next[i] = 0
		}
	}
	return next
}

// func strStr(haystack string, needle string) int {
// 	// 不借助工具的话要用KMP算法才能达到最优
// 	if len(needle) == 0 {
// 		return 0
// 	}
// 	for i := 0; i <= len(haystack)-len(needle); i++ {
// 		if haystack[i:i+len(needle)] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }

// @lc code=end

// func main(){
// 	strStr("mississippi","issip")
// }
