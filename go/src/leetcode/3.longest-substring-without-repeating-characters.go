/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package leetcode

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start, window := 0, 0

	for i, c := range s {
		if _, ok := charMap[c]; !ok || charMap[c] < start {
			window = int(math.Max(float64(i-start+1), float64(window)))
		} else {
			start = charMap[c] + 1
		}
		charMap[c] = i
	}
	return window
}

func lengthOfLongestSubstringAlt(s string) int {
	// 定义游标尺寸大小,游标的左边位置
	window, start := 0, 0

	// 循环字符串
	for key := 0; key < len(s); key++ {
		// 查看当前字符串是否在游标内
		isExist := strings.Index(string(s[start:key]), string(s[key]))

		// 如果不存在游标内部,游标长度重新计算并赋值
		if isExist == -1 {
			if key-start+1 > window {
				window = key - start + 1
			}
		} else { //存在，游标开始位置更换为重复字符串位置的下一个位置
			start = start + 1 + isExist
		}
	}
	return window
}

// func main() {
// 	s := "tmmzuxt"
// 	la := lengthOfLongestSubstringAlt(s)
// 	l := lengthOfLongestSubstring(s)
// 	fmt.Println(la)
// 	fmt.Println(l)
// }

// @lc code=end
