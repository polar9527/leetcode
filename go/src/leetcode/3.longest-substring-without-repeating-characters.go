/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	start, window := 0, 0

	for i, c := range s {
		if _, ok := charMap[c]; !ok || charMap[c] < start {
			charMap[c] = i
			window = int(math.Max(float64(i-start+1), float64(window)))
		} else {
			start = charMap[c] + 1
			charMap[c] = i
		}
	}
	return window
}

func lengthOfLongestSubstringAlt(s string) int {
	charMap := make(map[rune]int)
	start, window := 0, 0

	for i, c := range s {
		if _, ok := charMap[c]; !ok || charMap[c] < start {
			charMap[c] = i
			window = int(math.Max(float64(i-start+1), float64(window)))
		} else {
			start = charMap[c] + 1
			charMap[c] = i
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
