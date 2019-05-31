/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package main

// import (
// 	"strings"
// )

func isValid(b byte) bool {
	if ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	if s == "" || len(s) == 1 {
		return true
	}

	start, end := 0, len(s)-1
	v := strings.ToUpper(s)
	unStrCnt := 0

	for start < end {
		for start < len(s)-1 && !isValid(v[start]) {
			start++
			unStrCnt++
		}
		for end > 0 && !isValid(v[end]) {
			end--
			unStrCnt++
		}
		if v[start] == v[end] {
			start++
			end--
		} else {
			if unStrCnt == len(s) {
				return true
			} else {
				return false
			}
		}
	}
	return true

}

// func main() {
// 	s := ",a..b//"
// 	println(isPalindrome(s))
// }
