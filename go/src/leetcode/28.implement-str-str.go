package main

// package main

// package main

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

func strStr(haystack string, needle string) int {
	// 不借助工具的话要用KMP算法才能达到最优
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// func main() {
// strStr("mis", "mis")
// }
