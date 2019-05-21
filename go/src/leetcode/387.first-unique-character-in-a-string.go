/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

package main

func firstUniqChar(s string) int {
	var flags [26]int
	for i := range flags {
		flags[i] = -1
	}

	slen := len(s)
	for i, ch := range s {
		idx := byte(ch - 'a')
		if flags[idx] == -1 {
			flags[idx] = i
		} else {
			flags[idx] = slen
		}
	}
	min := slen
	for i := range flags {
		if flags[i] >= 0 && flags[i] < len(s) && flags[i] < min {
			min = flags[i]
		}
	}

	if min == slen {
		return -1
	}
	return min
}

// func main() {
// 	s := "adsfas"
// 	println(firstUniqChar(s))
// }
