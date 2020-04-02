package main

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	index := 0
	return isNumberCore(s, &index)

}

func isNumberCore(s string, index *int) bool {

	numeric := isInteger(s, index)

	if *index < len(s) && s[*index] == '.' {
		*index++
		numeric = isUnsigned(s, index) || numeric
	}

	if *index < len(s) && (s[*index] == 'e' || s[*index] == 'E') {
		*index++
		numeric = numeric && isInteger(s, index)
	}
	return numeric && len(s) == *index
}

func isInteger(s string, index *int) bool {
	if len(s[*index:]) == 0 {
		return false
	}
	if s[*index] == '+' || s[*index] == '-' {
		*index++
	}
	return isUnsigned(s, index)

}

func isUnsigned(s string, index *int) bool {
	if len(s[*index:]) == 0 {
		return false
	}
	start := *index
	for *index < len(s) {
		if s[*index] < '0' || s[*index] > '9' {
			break
		}
		*index++
	}
	return *index > start
}

// func main() {
// 	s := "124"
// 	isNumber(s)
//
// }
