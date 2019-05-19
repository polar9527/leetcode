/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

package main

func reverseString(s []byte) {
	size := len(s)
	h, t := 0, size-1

	for h < t {
		s[h], s[t] = s[t], s[h]
		h++
		t--
	}
	// println(h)
	// println(t)
	// println(string(s))
}

// func main() {
// 	s := "safsdf"
// 	sb := []byte(s)
// 	reverseString(sb)
// 	println(string(sb))
// }
