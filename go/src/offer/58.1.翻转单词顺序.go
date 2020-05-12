package main

import (
	"strings"
)

func reverseWords(s string) string {
	var reverseSeg []string
	seg := strings.Fields(s)

	for i := len(seg) - 1; i >= 0; i-- {
		reverseSeg = append(reverseSeg, seg[i])
	}
	return strings.Join(reverseSeg, " ")
}
