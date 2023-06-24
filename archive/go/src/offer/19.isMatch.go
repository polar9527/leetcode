package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		// if s is "" return true
		// if s is not "" return false
		return len(s) == 0
	}

	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		// next pattern
		ret1 := isMatch(s, p[2:])
		// next char
		ret2 := firstMatch && isMatch(s[1:], p)
		return ret1 || ret2
	} else {
		// 	len(p) >= 2 && p[1] != '*'
		ret := firstMatch && isMatch(s[1:], p[1:])
		return ret
	}
}

func main() {
	// ret := isMatch("aa", "a*")
	arr := []int{0, 1}

	fmt.Println(arr[2:])
}
